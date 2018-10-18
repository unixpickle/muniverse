(function () {

  const c2 = window.construct2api;
  let gameOver = false;

  window.muniverse = {
    init: function () {
      return c2.waitStart('Menu').then(function () {
        c2.runActions([
          '6850613773641528',
          '752169581943088',
          '4420041741113394',
          '3224401336031421',
          '6910949388048729',
          '2225224888692444',
          '9599451375703232',
          '4977243977992382',
          '5562208769801312',
          '119768221482659',
          '5310681519174606'
        ])
        window.faketime.pause();

        ['8551736235267092', '7578836874777334'].forEach((act) => {
          c2.hijackAction(act, () => gameOver = true);
        });

        // Remove restart and menu buttons.
        window.faketime.advance(100);
        ['210', '211'].forEach((x) => c2.disableObject(x));

        // Re-render the scene.
        window.faketime.advance(100);
      });
    },
    step: function (millis) {
      window.faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function () {
      return Promise.resolve(parseInt(cr_getC2Runtime().objectsByUid['67'].text));
    }
  };

})();