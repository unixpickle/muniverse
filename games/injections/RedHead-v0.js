(function() {

  var c2 = window.construct2api;

  var END_SCORE = '9526652099949740';
  var RUNNING_SCORE = '689137705713251';

  // Fix a page-loading bug due to a missing image.
  var old = window.fg_api.prototype.__;
  window.fg_api.prototype.__ = function(k) {
    if (k === 'more_games_image') {
      return window.fg_api.EMPTY_IMAGE;
    }
    return old.apply(this, arguments);
  }

  window.muniverse = {
    init: function() {
      return c2.waitStart('Menu').then(function() {
        cr_getC2Runtime().layouts.Menu.startRunning();
        cr_getC2Runtime().layouts.Level.startRunning();
        window.faketime.pause();
        window.faketime.advance(100);

        // Disable pause and refresh buttons
        [210, 211].forEach((x) => c2.disableObject(x));
        window.faketime.advance(100);
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(!!cr_getC2Runtime().actsBySid[END_SCORE].results[0]);
    },
    score: function() {
      return Promise.resolve(cr_getC2Runtime().actsBySid[END_SCORE].results[0] ||
        cr_getC2Runtime().actsBySid[RUNNING_SCORE].results[0] || 0);
    }
  };

})();
