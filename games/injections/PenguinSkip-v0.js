(function() {

  var OBJECT_UID_SCORE = 27;
  var OBJECT_UID_PAUSE = 117;
  var OBJECT_UID_GAME_OVER = 53;

  var c2 = window.construct2api;

  window.muniverse = {
    init: function() {
      return c2.waitStart('MainScreen').then(function() {
        var rt = cr_getC2Runtime();
        rt.layouts.MainScreen.stopRunning();
        rt.layouts.GameplayScreen.startRunning();
        rt.objectsByUid[OBJECT_UID_PAUSE].visible = false;

        // Skip past a frame of the main screen.
        window.faketime.advance(100);

        window.faketime.pause();
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      var obj = cr_getC2Runtime().objectsByUid[OBJECT_UID_GAME_OVER] || {};
      return Promise.resolve(!!(obj.layer || {}).visible);
    },
    score: function() {
      var obj = cr_getC2Runtime().objectsByUid[OBJECT_UID_SCORE];
      return Promise.resolve(obj.instance_vars[0]);
    }
  };

})();
