(function() {

  window.indexedDB.deleteDatabase('localforage');

  var c2 = window.construct2api;

  window.muniverse = {
    init: function() {
      return c2.waitStart('menu').then(function() {
        faketime.pause();
        var layouts = cr_getC2Runtime().layouts;
        layouts.menu.stopRunning();
        layouts.game.startRunning();
        faketime.advance(300);

        // Set to non-negative value on game over.
        c2.globalVar('highscore').setValue(-1);

        // Disable pause.
        c2.disableActions(["283391847588184"])

        // Disable restart.
        c2.disableActions([
          "698796394770349",
          "465513539272378",
          "815417831443698",
          "791168865452740",
          "542723129549946",
          "508451757460997",
          "301467848446384",
          "190100454728080",
          "526098579070463",
          "242857406401086",
          "442904037399484",
          "548813263529514",
          "234187344128948",
          "959632861215126",
          "119937900769678",
          "520610166685177",
          "487037849983345"
        ]);
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(c2.globalVar('highscore').getValue() >= 0);
    },
    score: function() {
      return Promise.resolve(c2.globalVar('score').getValue());
    }
  }

})();
