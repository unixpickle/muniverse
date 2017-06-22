(function() {

  var LOAD_TIMEOUT = 60000;

  var SLOGAN_SPRITE = [0, 2, 0, 7, 0];
  var SLOGAN = 'A Spot the Orange game';

  fg_api.EMPTY_IMAGE = 'dummy_image.png';

  localStorage.sto_fruitpulp_musicMuted = 'false';
  localStorage.sto_fruitpulp_sfxMuted = 'false';
  localStorage.sto_fruitpulp_state = JSON.stringify({
    strs: Array(36).fill(3),
  });

  function spriteAtPath(path) {
    var obj = ((window.scene || {}).root || {}).sprite;
    path.forEach((idx) => obj = ((obj || {}).children || [])[idx]);
    return obj || null;
  }

  var gameOver = false;
  var score = 0;

  window.muniverse = {
    init: function(options) {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return (spriteAtPath(SLOGAN_SPRITE) || {})._text === SLOGAN;
      }).then(function() {
        faketime.pause();

        // Disable all buttons from here on out.
        var oldB = ui.imageButton;
        ui.imageButton = function() {
          arguments[3] = () => false;
          return oldB.apply(this, arguments);
        };

        // Capture score changes.
        var oldMgr = scoring.manager;
        scoring.manager = function() {
          var a = oldMgr.call(scoring);
          a.scoreChanged.add((x) => score = x.updatedScore);
          return a;
        };

        var s = gamescene.create(null, options.level, null);
        scene.switchTo(s, scene.fade(0.1));

        ['gameovermenu', 'levelcompletedmenu'].forEach((menuName) => {
          var oldCreate = window[menuName].create;
          window[menuName].create = function() {
            gameOver = true;
            return oldCreate.apply(this, arguments);
          };
        });

        faketime.advance(2000);
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(score);
    }
  };

})();
