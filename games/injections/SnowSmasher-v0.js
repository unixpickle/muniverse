(function() {

  var c2 = window.construct2api;

  var gameOver = false;

  window.muniverse = {
    init: function() {
      return c2.waitStart('layoutGame').then(function() {
        window.faketime.pause();
        ['mousedown', 'mouseup'].forEach((evt) => {
          var e = new jQuery.Event(evt);
          e.pageX = document.body.offsetWidth / 2;
          e.pageY = document.body.offsetHeight / 2;
          jQuery(document).trigger(e);
        });

        var endgameLayout = cr_getC2Runtime().layouts.layoutResult;
        endgameLayout.startRunning = () => gameOver = true;

        window.faketime.advance(5000);
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(c2.globalVar('points').getValue());
    }
  };

})();
