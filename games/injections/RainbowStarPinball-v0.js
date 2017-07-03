(function() {

  var LOAD_TIMEOUT = 60000;

  // Serialization format:
  // http://old.haxe.org/manual/serialization/format.
  localStorage['B10RSP'] = 'oy12:_____VERSIONi1y16:defaultSessionIdoy9:'+
    'loadCountzy9:saveCounti1y9:highScorei0y16:isScrollDisabledtgg';

  var gameOver = false;
  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.globalMenu;
      }).then(function() {
        faketime.pause();

        // Skip initial animation.
        faketime.advance(1000);

        // Get rid of the pause button.
        globalMenu.view.parent.parent._children[1]._children.forEach((view) => {
          if (view.x === 275) {
            view.set_isVisible(false);
          }
        });

        globalMenu._pressPlay();
        faketime.advance(1000);

        if (!window.globalGame) {
          throw 'missing globalGame object';
        }
        globalGame._gameOver = () => gameOver = true;
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(globalGame._session.cache.score);
    }
  };

})();
