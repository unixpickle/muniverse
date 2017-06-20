(function() {

  // This game depends on a certain framerate.
  faketime.animationFrameRate = 60;

  window.indexedDB.deleteDatabase('localforage');

  var c2 = window.construct2api;

  window.muniverse = {
    init: function() {
      return c2.waitStart('Home').then(function() {
        window.faketime.pause();
        cr_getC2Runtime().layouts.Home.stopRunning();
        cr_getC2Runtime().layouts.Level1.startRunning();

        // Disable pause button.
        construct2api.disableActions(['3134628972584548'])

        window.faketime.advance(500);
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      var layout = cr_getC2Runtime().running_layout.name;
      return Promise.resolve(layout !== 'Level1');
    },
    score: function() {
      return Promise.resolve(c2.globalVar('Score').getValue());
    }
  };

})();
