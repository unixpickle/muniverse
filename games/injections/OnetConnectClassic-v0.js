(function() {

  var c2 = window.construct2api;

  // Prevent state from propagating between episodes.
  window.indexedDB.deleteDatabase('localforage');

  window.muniverse = {
    init: function() {
      return c2.waitStart('menu').then(function() {
        // Backup in case indexedDB doesn't work.
        window.localforage.clear();

        var layouts = cr_getC2Runtime().layouts;
        layouts.menu.stopRunning();
        layouts.game.startRunning();

        var pauseActions = [
          "225305688336864",
          "221372034106797",
          "993701224046407",
          "151089995842967",
          "123277322648489",
          "548280268767088",
          "918217655252041",
          "173931683769177",
          "458658086951332",
          "793239862526391",
          "530504484923984",
          "281471667038213",
          "772996049007964",
          "825741643870365",
          "920027563513016",
          "818730045083290",
          "262018315975087",
          "881103687496313"
        ];
        c2.disableActions(pauseActions);

        window.faketime.pause();

        // Re-render the scene already faded in.
        window.faketime.advance(750);
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);

      // Attempt at robust "game over" conditions:
      //
      // 1. Layer 4 has 3 instances if they clear a level.
      // 2. Level is not equal to one if they somehow manage
      //    to get onto the next level despite bullet 1.
      // 3. Layer 4 has 5 instances if the game ends.
      // 4. Timer is 0 if bullet 3 didn't work.
      // 5. We somehow got back to the main menu.
      //
      // Sometimes the game gets paused because of a popup
      // which uses one of our "shuffles".
      // This means that a move we made caused the game to
      // be unsolvable.
      // If we have no shuffles left, this will ultimately
      // lead to bullet 3.

      var layout = cr_getC2Runtime().running_layout;

      return Promise.resolve(layout.name !== 'game' ||
        layout.layers[4].instances.length === 3 ||
        layout.layers[4].instances.length === 5 ||
        c2.globalVar('level').getValue() !== 1 ||
        c2.globalVar('timer').getValue() === 0);
    },
    score: function() {
      return Promise.resolve(c2.globalVar('score').getValue());
    }
  };

  // For debugging, get unlimited tips:
  //
  //     window.construct2api.globalVar('tip').setValue(100);
  //
  // You can also force the game to end quickly:
  //
  //     window.construct2api.globalVar('timer').setValue(10);
  //
  // You can capture actions like so:
  //
  //     var actions = [];
  //     Object.keys(cr_getC2Runtime().actsBySid).forEach((x) => {
  //       var old = cr_getC2Runtime().actsBySid[x].run;
  //       cr_getC2Runtime().actsBySid[x].run = function() {
  //         if (actions.indexOf(x) < 0) actions.push(x);
  //         return old.apply(cr_getC2Runtime().actsBySid[x], arguments);
  //       }
  //     });
  //

})();
