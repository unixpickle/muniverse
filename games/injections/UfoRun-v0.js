(function() {

  var LOAD_TIMEOUT = 60000;

  // Found from obfuscated code.
  var ACTS_BY_SID = 'tj';
  var ACTION_RUN = 'zc';
  var VARIABLES = 'wy';
  var INITIAL_WIDTH = 'xc';
  var INITIAL_HEIGHT = 'wc';

  var gameOver = false;

  function replaceAction(act, func) {
    cr_getC2Runtime()[ACTS_BY_SID][act][ACTION_RUN] = func;
  };

  function playButtonCoordinates() {
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
    return [xAdd + 520*scale, yAdd + 320*scale];
  }

  function sendMouseEvent(evt, x, y) {
    var e = new jQuery.Event(evt);
    e.pageX = x;
    e.pageY = y;
    jQuery(document).trigger(e);
  }

  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.cr_getC2Runtime &&
          cr_getC2Runtime().Nj &&
          cr_getC2Runtime().Nj.loader &&
          !cr_getC2Runtime().Nj.loader.Uk &&
          cr_getC2Runtime().Nj.game &&
          cr_getC2Runtime().Nj.game.Ma &&
          cr_getC2Runtime().Nj.game.Ma.length === 3 &&
          cr_getC2Runtime().Nj.game.Ma[2].u &&
          cr_getC2Runtime().Nj.game.Ma[2].u.length === 18;
      }).then(function() {
        window.faketime.pause();

        // First action that seems to be triggered when
        // the game has fully ended.
        replaceAction('2680802799716757', () => gameOver = true);

        // Make sure main menu is fully visible.
        window.faketime.advance(1000);

        var coords = playButtonCoordinates();
        ['mousedown', 'mouseup'].forEach((e) => {
          sendMouseEvent(e, coords[0], coords[1]);
          window.faketime.advance(200);
        });

        var uiComps = cr_getC2Runtime().Nj.game.Ma[2].u;
        var objCount = uiComps.length;
        if (objCount !== 10) {
          throw 'unexpected object count: ' + objCount;
        }

        // Remove pause button.
        uiComps[7].visible = false;

        window.faketime.advance(100);
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(cr_getC2Runtime()[VARIABLES][16].data);
    }
  };

  // You can capture actions like so:
  //
  //     var actions = [];
  //     Object.keys(cr_getC2Runtime().tj).forEach((x) => {
  //       var old = cr_getC2Runtime().tj[x].zc;
  //       cr_getC2Runtime().tj[x].zc = function() {
  //         if (actions.indexOf(x) < 0) actions.push(x);
  //         return old.apply(cr_getC2Runtime().tj[x], arguments);
  //       }
  //     });
  //

})();
