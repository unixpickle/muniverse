(function() {

  var MENU_TIMEOUT = 60000;

  var PLAY_BUTTON = [0, 0, 1, 0];
  var CAR_BUTTONS = [
    [0, 6, 0], // VIPER
    [0, 7, 0]  // VICE
  ];
  var WEAPON_BUTTONS = [
    [0, 14, 0], // MACHINEGUN
    [0, 15, 0] // SHOTGUN
  ];
  var START_BUTTON = [0, 22, 0];

  // The game depends on 60fps animation frames.
  window.faketime.animationFrameRate = 60;

  function childAtPath(path) {
    var child = ((window.game || {}).stage) || {};
    path.forEach((x) => child = ((child || {}).children || [])[x]);
    return child || null;
  }

  function clickButton(button) {
    if (!button || !button.events || !button.events.onInputUp) {
      throw 'cannot click null button';
    }
    button.events.onInputUp.dispatch();
  }

  var gameOver = false;

  window.muniverse = {
    init: function() {
      var options = window.muniverse_variant || {car: 0, weapon: 0};
      return pollAndWait(MENU_TIMEOUT, function() {
        return 'undefined' !== typeof window.game &&
          window.game.isBooted &&
          window.game.isRunning &&
          childAtPath(PLAY_BUTTON);
      }).then(function() {
        // Disable pause button.
        States.Level.prototype.PauseLevel = function(){};
        States.Level.prototype.ShowResults = () => gameOver = true;

        // Use faketime to skip animations.
        window.faketime.pause();

        clickButton(childAtPath(PLAY_BUTTON));
        window.faketime.advance(1000);

        clickButton(childAtPath(CAR_BUTTONS[options.car]));
        window.faketime.advance(1000);

        clickButton(childAtPath(WEAPON_BUTTONS[options.weapon]));
        window.faketime.advance(1000);

        clickButton(childAtPath(START_BUTTON));
        window.game.paused = false;
        window.faketime.advance(100);
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(score);
    }
  };

})();
