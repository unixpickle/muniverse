(function() {

  var LOAD_TIMEOUT = 60000;
  var START_BUTTON = '6Nocu6swgqUvEgqppqUnrLfvuxLI2Q8BnTDgIcnqCVQ=';

  var oldUnderscore = fg_api.prototype.__;
  fg_api.prototype.getCurrentLanguage = () => 'en';
  fg_api.prototype.__ = function(k) {
    if (k === 'more_games_image') {
      // Data URL doesn't work.
      return 'library/spell/defaultAppearance.png';
    }
    return oldUnderscore.apply(this, arguments);
  };

  var gameOver = false;
  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        if (!window.require) {
          return false;
        }
        var spellObj = require('spell/client/main').spell;
        var scene = spellObj.sceneManager.activeScene;
        return scene.sceneConfig.name === 'GameScene';
      }).then(function() {
        window.faketime.pause();
        var spellObj = require('spell/client/main').spell;
        require(START_BUTTON).pointerDown(spellObj);
        window.faketime.advance(100);

        // Catch gameover signal as soon as it starts.
        var oldFunc = spellObj.entityManager.createEntity;
        spellObj.entityManager.createEntity = function(opts) {
          if (opts.name === 'gameover') {
            gameOver = true;
          }
          return oldFunc.apply(this, arguments);
        }
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      var spellObj = require('spell/client/main').spell;
      var entityMgr = spellObj.entityManager;
      var playerID = entityMgr.getEntityIdsByName('player')[0];
      var components = entityMgr.componentMaps;
      var component = components['kumbakarate.player.component'][playerID];
      return Promise.resolve(component.score);
    }
  };

})();
