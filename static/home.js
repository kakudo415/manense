'use strict';
var editHTML = document.getElementById('edit');
var editNameHTML = document.getElementById('editName');
var editIncomeHTML = document.getElementById('editIncome');
var editTimeHTML = document.getElementById('editTime');

var currentID;
function exMenu(exID, on) {
  currentID = exID;
  if (on) {
    var AJAX = new XMLHttpRequest();
    AJAX.open('POST', '/info', true);
    AJAX.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    AJAX.send('expense-uuid=' + exID);
    AJAX.onreadystatechange = () => {
      if (AJAX.readyState == 4 && AJAX.status == 200) {
        var data = JSON.parse(AJAX.responseText);
        editNameHTML.value = data.Name;
        editIncomeHTML.value = data.Income;
        editTimeHTML.value = data.Time.substr(0, 10);
        editHTML.style.display = 'block';
      }
    }
  } else {
    editHTML.style.display = 'none';
  }
}

function exSave() {
  var AJAX = new XMLHttpRequest();
  AJAX.open('POST', '/update', true);
  AJAX.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
  AJAX.send('expense-uuid=' + currentID + '&expense-name=' + editNameHTML.value + '&expense-income=' + editIncomeHTML.value + '&expense-time=' + editTimeHTML.value);
  AJAX.onreadystatechange = () => {
    if (AJAX.readyState == 4 && AJAX.status == 200) {
      exMenu('', false);
    }
  }
}

function exErase() {
  var AJAX = new XMLHttpRequest();
  AJAX.open('POST', '/erase', true);
  AJAX.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
  AJAX.send('expense-uuid=' + currentID);
  AJAX.onreadystatechange = () => {
    if (AJAX.readyState == 4 && AJAX.status == 200) {
      exMenu('', false);
      document.getElementById(currentID).style.display = "none";
    }
  }
}
