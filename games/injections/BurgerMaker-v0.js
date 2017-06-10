(function() {

  var LOAD_TIMEOUT = 60000;

  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return stage && stage.children.length === 1;
      }).then(function() {
        window.faketime.pause();
        goPage('game');
        window.faketime.advance(50);
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(stage.children[0].children[3].visible);
    },
    score: function() {
      return Promise.resolve(playerData.newScore);
    }
  };

})();
