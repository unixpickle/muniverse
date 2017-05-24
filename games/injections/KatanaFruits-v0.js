(function() {

  window.muniverse = {
    score: function() {
      return 0;
    },
    reset: function() {
      // TODO: navigate the menus and get to a game.
    },
    step: function() {
      // TODO: advance time here.
    }
  };

  var oldCInterface = window.CInterface;
  window.CInterface = function() {
    oldCInterface.call(this);

    var oldRefScore = this.refreshScore;
    this.refreshScore = function(x) {
      oldRefScore.call(this, x);
      window.muniverse.score = function() {
        return x;
      };
    }
  }

})();
