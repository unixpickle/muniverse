(function() {

  var LOAD_TIMEOUT = 60000;

  // Obfuscated keys.
  var INITIAL_WIDTH = 'qb';
  var INITIAL_HEIGHT = 'pb';

  // Coordinates of menu buttons.
  var PLAY_BUTTON = [288, 210];
  var CONTINUE_BUTTON = [443, 90];
  var CHARACTER_BUTTONS = [
    [200, 180],
    [370, 180]
  ];
  var GO_BUTTON = [470, 55];

  var gameOver = false;

  function replaceAction(act, func) {
    cr_getC2Runtime()[ACTS_BY_SID][act][ACTION_RUN] = func;
  };

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

  window.muniverse = {
    init: function() {
      var options = window.muniverse_variant || {character: 0};
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.cr_getC2Runtime &&
          cr_getC2Runtime().$ &&
          cr_getC2Runtime().$.name === 'Start';
      }).then(function() {
        window.faketime.pause();

        // First action that seems to be triggered on game over.
        cr_getC2Runtime().If['7228070045163318'].fb = () => gameOver = true;

        var buttons = [
          PLAY_BUTTON,
          CONTINUE_BUTTON,
          CHARACTER_BUTTONS[options.character],
          GO_BUTTON
        ];
        buttons.forEach((p) => {
          sendMouseEvent('mousedown', p[0], p[1]);
          sendMouseEvent('mouseup', p[0], p[1]);
          window.faketime.advance(50);
        });
        if (cr_getC2Runtime().$.name !== 'Layout 1') {
          throw 'unexpected layout after menu navigation';
        }
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      if (cr_getC2Runtime().$.name === 'GameOver' ||
          cr_getC2Runtime().$.name === 'Start') {
        gameOver = true;
      }
      return Promise.resolve(gameOver);
    },
    score: function() {
      return Promise.resolve(cr_getC2Runtime().Ef[3711933655180073].X[1] || 0);
    }
  };

  // You can capture actions like so:
  //
  //     var actions = [];
  //     Object.keys(cr_getC2Runtime().If).forEach((x) => {
  //       var old = cr_getC2Runtime().If[x].fb;
  //       cr_getC2Runtime().If[x].fb = function() {
  //         if (actions.indexOf(x) < 0) actions.push(x);
  //         return old.apply(cr_getC2Runtime().If[x], arguments);
  //       }
  //     });
  //
  // You can skip to the last level like so:
  //
  //     cr_getC2Runtime().bg['Start'].Iu();
  //     cr_getC2Runtime().bg['Layout 7'].bn();
  //

})();
