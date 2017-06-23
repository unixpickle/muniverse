(function() {

  var LOAD_TIMEOUT = 60000;

  var PLAY_BUTTON = [0, 54, 0];
  var LEVEL_BUTTONS = [0, 50, 1];
  var CHECK_BUTTON = [0, 48];

  // Avoid errors.
  _gaTrack = () => true;

  // Unlock all levels.
  localStorage['jwlxpl-data'] = JSON.stringify({
    snd: true,
    lng: 0,
    lfms: 0,
    lvls: Array(60).fill([0, 0])
  });

  function viewAtPath(path) {
    var obj = game.stage;
    path.forEach((idx) => obj = ((obj || {}).children || [])[idx]);
    return obj;
  }

  function clickButton(button) {
    if (!button) {
      throw 'cannot click null button';
    }
    button.clbck.call(button.clbckCtx, button);
  }

  function rawScore() {
    return r413.score;
  }

  var gameOver = false;
  var endgameScore = null;

  window.muniverse = {
    init: function(options) {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.game && viewAtPath(PLAY_BUTTON) &&
          window.r413 && window.r413.r372;
      }).then(function() {
        // Disable pause button.
        r413.r269 = () => false;
        r413.r372.clbck = () => false;

        r413.r9 = () => {
          gameOver = true;
          endgameScore = rawScore();
        };

        faketime.pause();
        clickButton(viewAtPath(PLAY_BUTTON));
        faketime.advance(1000);

        clickButton(viewAtPath(LEVEL_BUTTONS.concat([options.level])));
        faketime.advance(1000);

        var idx = 0;
        while (viewAtPath(CHECK_BUTTON).visible && idx < 10) {
          clickButton(viewAtPath(CHECK_BUTTON));
          faketime.advance(750);
          ++idx;
        }
        if (idx === 10) {
          throw 'check button not working';
        }
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(endgameScore === null ? rawScore() : endgameScore);
    }
  }

})();
