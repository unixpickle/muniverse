(function() {

  var LOAD_TIMEOUT = 60000;

  window.muniverse = {
    init: function() {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return window.cr_getC2Runtime &&
          cr_getC2Runtime().ga &&
          cr_getC2Runtime().ga.name === 'Menu';
      }).then(function() {
        faketime.pause();

        // Obfuscated stopRunning/startRunning combo.
        cr_getC2Runtime().il.Menu.Su()
        cr_getC2Runtime().il.Game.em();

        // Break the exit button.
        var exitActions = [
          "5631522997414485",
          "3131126497964036",
          "279367136988558",
          "1771488393619863",
          "2615059342655982",
          "7639227613415441",
          "3503806683417828",
          "969936123636564"
        ];
        var actsBySid = cr_getC2Runtime().Xd;
        exitActions.forEach((x) => actsBySid[x].Va = () => false);

        faketime.advance(300);
      });
    },
    step: function(millis) {
      faketime.advance(millis);
      return Promise.resolve(cr_getC2Runtime().hk[9].data !== 'game');
    },
    score: function() {
      return Promise.resolve(cr_getC2Runtime().hk[10].data);
    }
  }

  // Log actions:
  // var actions = [];
  // Object.keys(cr_getC2Runtime().Xd).forEach((x) => {
  //   var old = cr_getC2Runtime().Xd[x].Va;
  //   cr_getC2Runtime().Xd[x].Va = function() {
  //     if (actions.indexOf(x) < 0) actions.push(x);
  //     old.apply(cr_getC2Runtime().Xd[x], arguments);
  //   }
  // });

})();
