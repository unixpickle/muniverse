(function() {

  var c2 = window.construct2api;

  // Prevent state from propagating between episodes.
  window.indexedDB.deleteDatabase('localforage');

  var gameOver = false;
  window.muniverse = {
    init: function() {
      return c2.waitStart('Layout 1').then(function() {
        window.faketime.pause();

        ['mousedown', 'mouseup'].forEach((evtName) => {
          var e = new jQuery.Event(evtName);
          e.pageX = window.innerWidth / 2;
          e.pageY = window.innerHeight / 2;
          jQuery(document).trigger(e);
        });

        c2.hijackAction('461484439604306', () => gameOver = true);
        window.faketime.advance(1000);
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(c2.globalVar('score').getValue());
    }
  };

})();
