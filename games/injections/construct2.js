(function() {

  var LOAD_TIMEOUT = 60000;

  window.construct2api = {
    waitStart: function(name) {
      return pollAndWait(LOAD_TIMEOUT, function() {
        return typeof cr_getC2Runtime !== 'undefined' &&
          cr_getC2Runtime().running_layout &&
          cr_getC2Runtime().running_layout.name === name;
      });
    },
    globalVar: function(name) {
      var res;
      cr_getC2Runtime().all_global_vars.forEach((x) => {
        if (x.name === name) {
          res = x;
        }
      });
      return res;
    },
    runActions: function(actions) {
      actions.forEach((x) => cr_getC2Runtime().actsBySid[x].run());
    },
    disableActions: function(actions) {
      actions.forEach((x) => cr_getC2Runtime().actsBySid[x].run = () => null);
    },
    disableObject: function(uid) {
      var obj = cr_getC2Runtime().objectsByUid[uid] || {};
      obj.width = 0;
      obj.height = 0;
      obj.visible = false;
    }
  };

})();
