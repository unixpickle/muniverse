(function() {

  var c2 = window.construct2api;

  // Prevent image load error.
  var old__ = famobi.__;
  famobi.__ = function(key) {
    if (key === 'more_games_image') {
      return famobi.getMoreGamesButtonImage();
    }
    return old__.apply(this, arguments);
  };

  function playButtonX() {
    return document.body.offsetWidth / 2;
  }

  function playButtonY() {
    return document.body.offsetHeight * 0.4;
  }

  function sendPlayButtonEvent(evt) {
    var e = new jQuery.Event(evt);
    e.pageX = playButtonX();
    e.pageY = playButtonY();
    jQuery(document).trigger(e);
  }

  var gameOver = false;

  window.muniverse = {
    init: function(options) {
      return c2.waitStart('Game').then(function() {
        faketime.pause();
        faketime.advance(1000);
        c2.globalVar('Difficulty').setValue(options.difficulty);

        ['mousedown', 'mouseup'].map(sendPlayButtonEvent);
        faketime.advance(500);
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(!!localStorage.best);
    },
    score: function() {
      if (localStorage.best) {
        return Promise.resolve(parseInt(localStorage.best));
      } else {
        return Promise.resolve(c2.globalVar('Scores').getValue());
      }
    }
  };

})();
