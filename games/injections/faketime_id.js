// Extension to faketime.js which makes setInterval and
// setTimeout use the same identifier namespace.
// Technically this is not required, but some applications
// in the wild seem to depend on it.

(function() {
  var _oldSetInterval = window.setInterval;
  var _oldClearTimeout = window.clearTimeout;
  var _oldClearInterval = window.clearInterval;

  window.setInterval = function() {
    return -_oldSetInterval.apply(this, arguments);
  };

  window.clearTimeout = window.clearInterval = function(id) {
    if (id > 0) {
      _oldClearTimeout(id);
    } else {
      _oldClearInterval(-id);
    }
  };
})();
