// Stubs to make Famobi games work correctly.

var _fgq = [];

function fg_api(data) {
  this._data = data;
}

fg_api.prototype.__ = function(key) {
  var res = this._data.game_i18n.en[key];
  return res.replace('{lang}', 'en');
}

fg_api.prototype.getMoreGamesButtonImage = function() {
  return '';
};

fg_api.prototype.moreGamesLink = function() {
  return '';
};

fg_api.prototype.submitHighscore = function() {
};
