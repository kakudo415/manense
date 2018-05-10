'use strict';
var balanceHTML = document.getElementById('balance');
var editHTML = document.getElementById('edit');
var editNameHTML = document.getElementById('editName');
var editIncomeHTML = document.getElementById('editIncome');
var editTimeHTML = document.getElementById('editTime');

var currentID;
function exMenu(exID, on) {
  currentID = exID;
  if (on) {
    if (currentID.length > 0) {
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
      editHTML.style.display = 'block';
    }
  } else {
    editHTML.style.display = 'none';
  }
}

function exSave() {
  var AJAX = new XMLHttpRequest();
  if (currentID.length > 0) {
    AJAX.open('POST', '/update', true);
    AJAX.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    AJAX.send('expense-uuid=' + currentID + '&expense-name=' + editNameHTML.value + '&expense-income=' + editIncomeHTML.value + '&expense-time=' + editTimeHTML.value);
    AJAX.onreadystatechange = () => {
      if (AJAX.readyState == 4 && AJAX.status == 200) {
        document.getElementById(currentID).innerHTML = `<span>${editNameHTML.value}</span><span>${editTimeHTML.value}</span><span>${editIncomeHTML.value}</span>`;
        balanceHTML.innerText = "残高 : " + AJAX.responseText;
        exMenu('', false);
      }
    }
  } else {
    AJAX.open('POST', '/new', true);
    AJAX.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    AJAX.send('expense-name=' + editNameHTML.value + '&expense-income=' + editIncomeHTML.value + '&expense-time=' + editTimeHTML.value);
    AJAX.onreadystatechange = () => {
      if (AJAX.readyState == 4 && AJAX.status == 200) {
        var data = JSON.parse(AJAX.responseText);
        document.getElementById('exs').innerHTML = `<a id=${data.uuid} class="ex" onclick="exMenu('${data.uuid}', true);"><span>${editNameHTML.value}</span><span>${editTimeHTML.value}</span><span>${editIncomeHTML.value}</span></a>` + document.getElementById('exs').innerHTML;
        balanceHTML.innerText = "残高 : " + data.balance;
        exMenu('', false);
      }
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
      document.getElementById(currentID).style.display = 'none';
      console.log(AJAX.responseText);
      balanceHTML.innerText = "残高 : " + AJAX.responseText;
      exMenu('', false);
    }
  }
}
