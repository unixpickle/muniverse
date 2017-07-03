(function() {

  var c2 = window.construct2api;

  window.indexedDB.deleteDatabase('localforage');

  var gameOver = false;

  window.muniverse = {
    init: function() {
      return c2.waitStart('menu').then(function() {
        faketime.pause();
        cr_getC2Runtime().layouts.menu.stopRunning();
        cr_getC2Runtime().layouts.level.startRunning();
        faketime.advance(50);

        construct2api.hijackAction('8839194560944502', () => gameOver = true);
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(c2.globalVar('coins').getValue());
    }
  };

})();
