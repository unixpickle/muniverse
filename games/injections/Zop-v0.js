(function() {

  var LOAD_TIMEOUT = 60000;

  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return 'number' === typeof window.time &&
          document.getElementsByClassName('z-footer').length === 1;
      }).then(function() {
        document.getElementsByClassName('z-footer')[0].style.display = 'none';
        window.gameRestart = () => false;
        window.faketime.pause();
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(window.time <= 0);
    },
    score: function() {
      return Promise.resolve(window.score);
    }
  };

})();
