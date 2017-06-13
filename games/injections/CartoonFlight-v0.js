(function() {

  var INITIAL_WIDTH = 'cb';
  var INITIAL_HEIGHT = 'zc';

  var LOAD_TIMEOUT = 60000;

  // Prevent 404 error.
  var old__ = window.fg_api.prototype.__;
  window.fg_api.prototype.__ = function(k) {
    if (k === 'more_games_image') {
      return 'icon-256.png';
    }
    return old__.apply(this, arguments);
  };

  // Prevent special treatment on first game.
  localStorage['H_best_distance'] = '1';

  function sendMouseEvent(evt, x, y) {
    var width = window.innerWidth;
    var height = window.innerHeight;
    var rt = cr_getC2Runtime();
    var origWidth = rt[INITIAL_WIDTH];
    var origHeight = rt[INITIAL_HEIGHT];
    var rt = cr_getC2Runtime();
    var xScale = width / origWidth;
    var yScale = height / origHeight;
    var scale = Math.min(xScale, yScale);
    var xAdd = (width - scale*origWidth) / 2;
    var yAdd = (height - scale*origHeight) / 2;

    var e = new jQuery.Event(evt);
    e.pageX = xAdd + x*scale;
    e.pageY = y*scale;
    jQuery(document).trigger(e);
  }

  var gameOver = false;
  var scoreMode = '';

  window.muniverse = {
    init: function(options) {
      scoreMode = options.scoreMode;
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.cr_getC2Runtime &&
          cr_getC2Runtime().qa &&
          cr_getC2Runtime().qa.name === 'menu';
      }).then(function() {
        window.faketime.pause();

        var pauseActions = [
          "7588007547677455",
          "1305298150821965",
          "4036630583934544",
          "2469973274296541",
          "920292507265257",
          "6038087539532941",
          "2092608481681369",
          "9892887063881912",
          "2389531679837251",
          "9864852929657058",
          "7148298800157745",
          "5331755548504835",
          "5160334450934046",
          "1876613047217508",
          "6300197537509727",
          "1735019484003593",
          "9228375042859862",
          "8111755648696253"
        ];
        pauseActions.forEach((x) => cr_getC2Runtime().xj[x].Dc = function(){});

        cr_getC2Runtime().xj['9392300557340988'].Dc = () => gameOver = true;

        ['mousedown', 'mouseup'].forEach(function(evt) {
          sendMouseEvent(evt, 350, 573);
        });
        window.faketime.advance(500);
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      if (scoreMode === 'stars') {
        return Promise.resolve(cr_getC2Runtime().Ex[12].data);
      } else if (scoreMode === 'distance') {
        return Promise.resolve(cr_getC2Runtime().Ex[15].data);
      } else {
        return Promise.reject('bad scoreMode: ' + scoreMode);
      }
    }
  };

  // You can capture actions like so:
  //
  //     var actions = [];
  //     Object.keys(cr_getC2Runtime().xj).forEach((x) => {
  //       var old = cr_getC2Runtime().xj[x].Dc;
  //       cr_getC2Runtime().xj[x].Dc = function() {
  //         if (actions.indexOf(x) < 0) actions.push(x);
  //         return old.apply(cr_getC2Runtime().xj[x], arguments);
  //       }
  //     });
  //

})();
