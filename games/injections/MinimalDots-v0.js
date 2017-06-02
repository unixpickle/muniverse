(function() {

  var MENU_TIMEOUT = 60000;

  function getGlobalVar(name) {
    var res;
    cr_getC2Runtime().all_global_vars.forEach((x) => x.name === name && (res = x));
    return res;
  }

  window.muniverse = {
    init: function() {
      return pollAndWait(MENU_TIMEOUT, function() {
        return typeof cr_getC2Runtime !== 'undefined' &&
          cr_getC2Runtime().running_layout &&
          cr_getC2Runtime().running_layout.name === 'Game';
      }).then(function() {
        var actions = [
          "8367457620323879",
          "3273814180686107",
          "9741637885575604",
          "6970904565799639",
          "5752549763901044",
          "7886278298902863",
          "8446293208042638",
          "4899538162202658",
          "8539905980676693",
          "2403910773354043",
          "8790829122107103",
          "2334484803472091",
          "8261822201394956",
          "3576332447348401",
          "7930276117275005",
          "5396138479594926",
          "9695276239857680",
          "3349008999467318",
          "5793654983654301",
          "8561624533709254",
          "5962084291111393",
          "2274777833210239",
          "6218497423232376",
          "2782627431056424",
          "2480756798654173",
          "5739175334508304",
          "8838488641294781",
          "6881488625941642",
          "9063140082916856",
          "9196598953632876",
          "3365321696316076",
          "9618512613392364"
        ];
        actions.forEach((x) => cr_getC2Runtime().actsBySid[x].run());
        window.faketime.pause();

        // Re-render the scene.
        window.faketime.advance(100);
      });
    },
    step: function(millis) {
      window.faketime.advance(millis);
      return Promise.resolve(getGlobalVar('GameOver').getValue() === 1);
    },
    score: function() {
      return Promise.resolve(getGlobalVar('Score').getValue());
    }
  };

})();
