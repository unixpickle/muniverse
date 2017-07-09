(function() {

  var LOAD_TIMEOUT = 60000;

  localStorage.clear();

  var gameOver = false;
  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.cr_getC2Runtime &&
          cr_getC2Runtime().la &&
          cr_getC2Runtime().la.name === 'menu';
      }).then(function() {
        faketime.pause();

        // Hit play button.
        ['mousedown', 'mouseup'].forEach((evt) => {
          var e = new jQuery.Event(evt);
          e.pageX = document.body.offsetWidth / 2;
          e.pageY = document.body.offsetHeight / 2;
          jQuery(document).trigger(e);
        });
        faketime.advance(100);

        // Disable home button (and all other buttons).
        jQuery(document).off('mousedown mouseup');

        cr_getC2Runtime().Se[1455297146827796].Za = () => gameOver = true;
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(parseInt(localStorage.score) || 0);
    }
  };

})();
