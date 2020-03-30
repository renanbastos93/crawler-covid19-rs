package main

import (
	"reflect"
	"testing"
)

var body string

func mockBody() {
	body = `<html><head>
	<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css" integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous">
	<script src="https://code.jquery.com/jquery-3.4.1.slim.min.js" integrity="sha384-J6qa4849blE2+poT4WnyKhv5vZF5SrPo0iEjwBvKU7imGFAV0wwj1yYfoRSJoZ+n" crossorigin="anonymous"></script><script data-dapp-detection="">
	(function() {
	  let alreadyInsertedMetaTag = false
	
	  function __insertDappDetected() {
		if (!alreadyInsertedMetaTag) {
		  const meta = document.createElement('meta')
		  meta.name = 'dapp-detected'
		  document.head.appendChild(meta)
		  alreadyInsertedMetaTag = true
		}
	  }
	
	  if (window.hasOwnProperty('web3')) {
		// Note a closure can't be used for this var because some sites like
		// www.wnyc.org do a second script execution via eval for some reason.
		window.__disableDappDetectionInsertion = true
		// Likely oldWeb3 is undefined and it has a property only because
		// we defined it. Some sites like wnyc.org are evaling all scripts
		// that exist again, so this is protection against multiple calls.
		if (window.web3 === undefined) {
		  return
		}
		__insertDappDetected()
	  } else {
		var oldWeb3 = window.web3
		Object.defineProperty(window, 'web3', {
		  configurable: true,
		  set: function (val) {
			if (!window.__disableDappDetectionInsertion)
			  __insertDappDetected()
			oldWeb3 = val
		  },
		  get: function () {
			if (!window.__disableDappDetectionInsertion)
			  __insertDappDetected()
			return oldWeb3
		  }
		})
	  }
	})()</script>
	<script src="https://cdn.jsdelivr.net/npm/popper.js@1.16.0/dist/umd/popper.min.js" integrity="sha384-Q6E9RHvbIyZFJoft+2mJbHaEWldlvI9IOYy5n3zV9zzTtmI3UksdQRVvoxMfooAo" crossorigin="anonymous"></script>
	<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/js/bootstrap.min.js" integrity="sha384-wfSDF2E50Y2D1uUdj0O3uMBJnjuUD4Ih7YwaYd1iqfktj0Uod8GCExl3Og8ifwB6" crossorigin="anonymous"></script>
	<script src="https://www.chartjs.org/dist/2.9.3/Chart.min.js"></script>
	<script src="https://www.chartjs.org/utils.js"></script>
	<link rel="stylesheet" type="text/css" href="http://cdn.leafletjs.com/leaflet/v0.7.7/leaflet.css">
	<link rel="stylesheet" type="text/css" href="http://cdnjs.cloudflare.com/ajax/libs/leaflet.markercluster/0.4.0/MarkerCluster.css">
	<link rel="stylesheet" type="text/css" href="http://cdnjs.cloudflare.com/ajax/libs/leaflet.markercluster/0.4.0/MarkerCluster.Default.css">
	<script type="text/javascript" src="http://ajax.googleapis.com/ajax/libs/jquery/2.1.4/jquery.min.js"></script>
	<script type="text/javascript" src="http://cdn.leafletjs.com/leaflet/v0.7.7/leaflet.js"></script>
	<script type="text/javascript" src="http://cdnjs.cloudflare.com/ajax/libs/leaflet.markercluster/0.4.0/leaflet.markercluster.js"></script>
	<title>SES-RS :: Covid19</title>
	<!-- Global site tag (gtag.js) - Google Analytics -->
	<script async="" src="https://www.googletagmanager.com/gtag/js?id=UA-161782390-1"></script>
	<script>
	  window.dataLayer = window.dataLayer || [];
	  function gtag(){dataLayer.push(arguments);}
	  gtag('js', new Date());
	  gtag('config', 'UA-161782390-1');
	</script>
	<style>
	canvas { -moz-user-select: none; -webkit-user-select: none;	-ms-user-select: none;}
	.map-responsive{    overflow:hidden;    padding-bottom:80%;    position:relative;    height:0; border: 1px solid #666666}
	button { background-color:#ffffff;border:1px solid #666666; font-size:10;font-family:Verdana; padding-left:10px; padding-right:10px; padding-top:3px; padding-bottom:3px }
	
	.ClassGraph { height: 730px; }
	
	</style>
	<style type="text/css">/* Chart.js */
	@keyframes chartjs-render-animation{from{opacity:.99}to{opacity:1}}.chartjs-render-monitor{animation:chartjs-render-animation 1ms}.chartjs-size-monitor,.chartjs-size-monitor-expand,.chartjs-size-monitor-shrink{position:absolute;direction:ltr;left:0;top:0;right:0;bottom:0;overflow:hidden;pointer-events:none;visibility:hidden;z-index:-1}.chartjs-size-monitor-expand>div{position:absolute;width:1000000px;height:1000000px;left:0;top:0}.chartjs-size-monitor-shrink>div{position:absolute;width:200%;height:200%;left:0;top:0}</style></head>
	<body data-gr-c-s-loaded="true">
	<div class="container-fluid">
		<div class="row">
			<div class="col col-xs-12"><img src="coronavirus.png"></div>
			<div class="col col-xs-12 text-center align-self-center"><h3>239</h3> <h5>confirmados</h5></div>
			<!--<div class="col col-xs-12 text-center align-self-center"><h3>57</h3> <h5>confirmados hoje</h5></div>-->
			<div class="col col-xs-12 text-center align-self-center"><h3>2</h3> <h5>mortes no total</h5></div>
		</div>
		<div class="row ClassGraph">
			<div class="col-md-3 col-xs-12"><div class="chartjs-size-monitor"><div class="chartjs-size-monitor-expand"><div class=""></div></div><div class="chartjs-size-monitor-shrink"><div class=""></div></div></div><canvas id="GraphMunicipio" style="display: block; width: 479px; height: 150px;" width="479" height="150" class="chartjs-render-monitor"></canvas></div>
			<div class="col-md-9 col-xs-12">
				<div class="row">
					<div class="col-md-8 col-xs-12">
						<button id="btnConf" onclick="javascript:trocamapa(1)">Casos Confirmados</button><button id="btnObitos" onclick="javascript:trocamapa(2,map,mapObitos)" style="background-color:#dfdfdf">Apenas Óbitos</button>
						<div class="map-responsive leaflet-container leaflet-fade-anim" id="map" name="map" tabindex="0"><div class="leaflet-map-pane" style="transform: translate3d(0px, 0px, 0px);"><div class="leaflet-tile-pane"><div class="leaflet-layer"><div class="leaflet-tile-container"></div><div class="leaflet-tile-container leaflet-zoom-animated"><img class="leaflet-tile leaflet-tile-loaded" src="http://c.tile.openstreetmap.org/6/22/37.png" style="height: 256px; width: 256px; left: 91px; top: 13px;"><img class="leaflet-tile leaflet-tile-loaded" src="http://b.tile.openstreetmap.org/6/22/36.png" style="height: 256px; width: 256px; left: 91px; top: -243px;"><img class="leaflet-tile leaflet-tile-loaded" src="http://b.tile.openstreetmap.org/6/21/37.png" style="height: 256px; width: 256px; left: -165px; top: 13px;"><img class="leaflet-tile leaflet-tile-loaded" src="http://a.tile.openstreetmap.org/6/23/37.png" style="height: 256px; width: 256px; left: 347px; top: 13px;"><img class="leaflet-tile leaflet-tile-loaded" src="http://a.tile.openstreetmap.org/6/22/38.png" style="height: 256px; width: 256px; left: 91px; top: 269px;"><img class="leaflet-tile leaflet-tile-loaded" src="http://a.tile.openstreetmap.org/6/21/36.png" style="height: 256px; width: 256px; left: -165px; top: -243px;"><img class="leaflet-tile leaflet-tile-loaded" src="http://c.tile.openstreetmap.org/6/23/36.png" style="height: 256px; width: 256px; left: 347px; top: -243px;"><img class="leaflet-tile leaflet-tile-loaded" src="http://c.tile.openstreetmap.org/6/21/38.png" style="height: 256px; width: 256px; left: -165px; top: 269px;"><img class="leaflet-tile leaflet-tile-loaded" src="http://b.tile.openstreetmap.org/6/23/38.png" style="height: 256px; width: 256px; left: 347px; top: 269px;"></div></div></div><div class="leaflet-objects-pane"><div class="leaflet-shadow-pane"></div><div class="leaflet-overlay-pane"><svg class="leaflet-zoom-animated" width="663" height="532" viewBox="-93 -74 663 532" style="transform: translate3d(-93px, -74px, 0px);"></svg></div><div class="leaflet-marker-pane"><img src="http://ti.saude.rs.gov.br/covid19/leaflet/images/marker.png" class="leaflet-marker-icon leaflet-zoom-animated leaflet-clickable" tabindex="0" style="margin-left: -9px; margin-top: -21px; width: 25px; height: 41px; transform: translate3d(172px, 55px, 0px); z-index: 55;"><img src="http://ti.saude.rs.gov.br/covid19/leaflet/images/marker.png" class="leaflet-marker-icon leaflet-zoom-animated leaflet-clickable" tabindex="0" style="margin-left: -9px; margin-top: -21px; width: 25px; height: 41px; transform: translate3d(154px, 123px, 0px); z-index: 123;"><img src="http://ti.saude.rs.gov.br/covid19/leaflet/images/marker.png" class="leaflet-marker-icon leaflet-zoom-animated leaflet-clickable" tabindex="0" style="margin-left: -9px; margin-top: -21px; width: 25px; height: 41px; transform: translate3d(244px, 167px, 0px); z-index: 167;"><div class="leaflet-marker-icon marker-cluster marker-cluster-small leaflet-zoom-animated leaflet-clickable" tabindex="0" style="margin-left: -20px; margin-top: -20px; width: 40px; height: 40px; transform: translate3d(207px, 145px, 0px); z-index: 145;"><div><span>3</span></div></div><div class="leaflet-marker-icon marker-cluster marker-cluster-small leaflet-zoom-animated leaflet-clickable" tabindex="0" style="margin-left: -20px; margin-top: -20px; width: 40px; height: 40px; transform: translate3d(244px, 241px, 0px); z-index: 241;"><div><span>2</span></div></div><div class="leaflet-marker-icon marker-cluster marker-cluster-small leaflet-zoom-animated leaflet-clickable" tabindex="0" style="margin-left: -20px; margin-top: -20px; width: 40px; height: 40px; transform: translate3d(274px, 266px, 0px); z-index: 266;"><div><span>4</span></div></div><div class="leaflet-marker-icon marker-cluster marker-cluster-small leaflet-zoom-animated leaflet-clickable" tabindex="0" style="margin-left: -20px; margin-top: -20px; width: 40px; height: 40px; transform: translate3d(272px, 43px, 0px); z-index: 43;"><div><span>3</span></div></div><div class="leaflet-marker-icon marker-cluster marker-cluster-small leaflet-zoom-animated leaflet-clickable" tabindex="0" style="margin-left: -20px; margin-top: -20px; width: 40px; height: 40px; transform: translate3d(122px, 204px, 0px); z-index: 204;"><div><span>4</span></div></div><img src="http://ti.saude.rs.gov.br/covid19/leaflet/images/marker.png" class="leaflet-marker-icon leaflet-zoom-animated leaflet-clickable" tabindex="0" style="margin-left: -9px; margin-top: -21px; width: 25px; height: 41px; transform: translate3d(163px, 217px, 0px); z-index: 217;"><div class="leaflet-marker-icon marker-cluster marker-cluster-small leaflet-zoom-animated leaflet-clickable" tabindex="0" style="margin-left: -20px; margin-top: -20px; width: 40px; height: 40px; transform: translate3d(189px, 236px, 0px); z-index: 236;"><div><span>9</span></div></div><img src="http://ti.saude.rs.gov.br/covid19/leaflet/images/marker.png" class="leaflet-marker-icon leaflet-zoom-animated leaflet-clickable" tabindex="0" style="margin-left: -9px; margin-top: -21px; width: 25px; height: 41px; transform: translate3d(266px, 75px, 0px); z-index: 75;"><div class="leaflet-marker-icon marker-cluster marker-cluster-small leaflet-zoom-animated leaflet-clickable" tabindex="0" style="margin-left: -20px; margin-top: -20px; width: 40px; height: 40px; transform: translate3d(286px, 99px, 0px); z-index: 99;"><div><span>7</span></div></div><div class="leaflet-marker-icon marker-cluster marker-cluster-small leaflet-zoom-animated leaflet-clickable" tabindex="0" style="margin-left: -20px; margin-top: -20px; width: 40px; height: 40px; transform: translate3d(286px, 138px, 0px); z-index: 138;"><div><span>7</span></div></div><div class="leaflet-marker-icon marker-cluster marker-cluster-medium leaflet-zoom-animated leaflet-clickable" tabindex="0" style="margin-left: -20px; margin-top: -20px; width: 40px; height: 40px; transform: translate3d(312px, 123px, 0px); z-index: 123;"><div><span>11</span></div></div><div class="leaflet-marker-icon marker-cluster marker-cluster-small leaflet-zoom-animated leaflet-clickable" tabindex="0" style="margin-left: -20px; margin-top: -20px; width: 40px; height: 40px; transform: translate3d(371px, 155px, 0px); z-index: 155;"><div><span>3</span></div></div><div class="leaflet-marker-icon marker-cluster marker-cluster-small leaflet-zoom-animated leaflet-clickable" tabindex="0" style="margin-left: -20px; margin-top: -20px; width: 40px; height: 40px; transform: translate3d(388px, 130px, 0px); z-index: 130;"><div><span>7</span></div></div><img src="http://ti.saude.rs.gov.br/covid19/leaflet/images/marker.png" class="leaflet-marker-icon leaflet-zoom-animated leaflet-clickable" tabindex="0" style="margin-left: -9px; margin-top: -21px; width: 25px; height: 41px; transform: translate3d(296px, 196px, 0px); z-index: 196;"><div class="leaflet-marker-icon marker-cluster marker-cluster-small leaflet-zoom-animated leaflet-clickable" tabindex="0" style="margin-left: -20px; margin-top: -20px; width: 40px; height: 40px; transform: translate3d(347px, 150px, 0px); z-index: 150;"><div><span>3</span></div></div><div class="leaflet-marker-icon marker-cluster marker-cluster-medium leaflet-zoom-animated leaflet-clickable" tabindex="0" style="margin-left: -20px; margin-top: -20px; width: 40px; height: 40px; transform: translate3d(325px, 148px, 0px); z-index: 148;"><div><span>13</span></div></div><div class="leaflet-marker-icon marker-cluster marker-cluster-large leaflet-zoom-animated leaflet-clickable" tabindex="0" style="margin-left: -20px; margin-top: -20px; width: 40px; height: 40px; transform: translate3d(321px, 164px, 0px); z-index: 164;"><div><span>157</span></div></div></div><div class="leaflet-popup-pane"></div></div></div><div class="leaflet-control-container"><div class="leaflet-top leaflet-left"><div class="leaflet-control-zoom leaflet-bar leaflet-control"><a class="leaflet-control-zoom-in" href="#" title="Zoom in">+</a><a class="leaflet-control-zoom-out leaflet-disabled" href="#" title="Zoom out">-</a></div></div><div class="leaflet-top leaflet-right"></div><div class="leaflet-bottom leaflet-left"></div><div class="leaflet-bottom leaflet-right"><div class="leaflet-control-attribution leaflet-control"><a href="http://leafletjs.com" title="A JS library for interactive maps">Leaflet</a> | © <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> | SES-RS/CEVS/DGTI | Atualizado em 29/03/2020 20:18:12</div></div></div></div>
						<div class="map-responsive leaflet-container leaflet-fade-anim" id="mapObitos" name="mapObitos" style="display:none;" tabindex="0"><div class="leaflet-map-pane" style="transform: translate3d(0px, 0px, 0px);"><div class="leaflet-tile-pane"><div class="leaflet-layer"><div class="leaflet-tile-container"></div><div class="leaflet-tile-container leaflet-zoom-animated"><img class="leaflet-tile leaflet-tile-loaded" src="http://c.tile.openstreetmap.org/6/22/37.png" style="height: 256px; width: 256px; left: -148px; top: -179px;"></div></div></div><div class="leaflet-objects-pane"><div class="leaflet-shadow-pane"></div><div class="leaflet-overlay-pane"><svg class="leaflet-zoom-animated" width="0" height="0" viewBox="0 0 0 0" style="transform: translate3d(0px, 0px, 0px);"></svg></div><div class="leaflet-marker-pane"></div><div class="leaflet-popup-pane"></div></div></div><div class="leaflet-control-container"><div class="leaflet-top leaflet-left"><div class="leaflet-control-zoom leaflet-bar leaflet-control"><a class="leaflet-control-zoom-in" href="#" title="Zoom in">+</a><a class="leaflet-control-zoom-out leaflet-disabled" href="#" title="Zoom out">-</a></div></div><div class="leaflet-top leaflet-right"></div><div class="leaflet-bottom leaflet-left"></div><div class="leaflet-bottom leaflet-right"><div class="leaflet-control-attribution leaflet-control"><a href="http://leafletjs.com" title="A JS library for interactive maps">Leaflet</a> | © <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> | SES-RS/CEVS/DGTI | Atualizado em 29/03/2020 20:18:12</div></div></div></div>
					</div>
					<div class="col-md-4 col-xs-12"><div class="chartjs-size-monitor"><div class="chartjs-size-monitor-expand"><div class=""></div></div><div class="chartjs-size-monitor-shrink"><div class=""></div></div></div><canvas id="GraphIdade" width="479" height="150" class="chartjs-render-monitor" style="display: block; width: 479px; height: 150px;"></canvas></div>
				</div>
				<div class="row">
					<div class="col-md-3"><div class="chartjs-size-monitor"><div class="chartjs-size-monitor-expand"><div class=""></div></div><div class="chartjs-size-monitor-shrink"><div class=""></div></div></div>
						<canvas id="GraphSexo" width="479" height="150" class="chartjs-render-monitor" style="display: block; width: 479px; height: 150px;"></canvas>
					</div>
					<div class="col-md-9"><div class="chartjs-size-monitor"><div class="chartjs-size-monitor-expand"><div class=""></div></div><div class="chartjs-size-monitor-shrink"><div class=""></div></div></div>
						<canvas id="GraphTotalCasos" style="position: relative; height: 201px; display: block; width: 479px;" width="479" height="201" class="chartjs-render-monitor"></canvas>
					</div>
				</div>
			</div>
		</div>
	</div>
	<script>
	var ctx = document.getElementById("GraphMunicipio").getContext('2d');
	var myChart = new Chart(ctx, {type: 'horizontalBar',data: {labels: [" Porto Alegre", " Bagé", " Canoas"],datasets: [{data: [143,9,8],backgroundColor: 'rgba(166, 58, 58, 0.8)',borderColor: 'rgba(0,0,0,0.8)',borderWidth: 1}]},
	options: {
	"hover": {"animationDuration": 0},"animation": { "duration": 1,"onComplete": function () { var chartInstance = this.chart, ctx = chartInstance.ctx; ctx.font = Chart.helpers.fontString(Chart.defaults.global.defaultFontSize-1, Chart.defaults.global.defaultFontStyle, Chart.defaults.global.defaultFontFamily); ctx.textAlign = 'center'; ctx.textBaseline = 'bottom'; this.data.datasets.forEach(function (dataset, i) { var meta = chartInstance.controller.getDatasetMeta(i); meta.data.forEach(function (bar, index) { var data = dataset.data[index]; ctx.fillText(data, bar._model.x + 10, bar._model.y + 7); });});}},
	maintainAspectRatio: false,legend: {display: false}, title: {display: true,text: ["Município de Residência","Total de Municípios: 47"]},scales: {yAxes: [{ticks: {beginAtZero: true}}]}}});
	
	/*
	var ctxB = document.getElementById("GraphConfirmados").getContext('2d');
	var myBarChart = new Chart(ctxB, {
	type: 'bar',data: {	labels: ["29/02", "01/03", "09/03", "11/03", "12/03", "13/03", "14/03", "15/03", "16/03", "17/03", "18/03", "19/03", "20/03", "21/03", "22/03", "23/03", "24/03", "25/03", "26/03", "27/03"],	datasets: [{	label: 'confirmados',	data: [1,1,3,2,4,5,10,7,29,20,26,24,29,7,12,17,20,9,6,7],	backgroundColor: 'rgba(166, 58, 58, 0.8)',	borderColor: 'rgba(0,0,0,1)',	borderWidth: 1	}]},options: {
	"hover": {"animationDuration": 0},"animation": { "duration": 1,"onComplete": function () { var chartInstance = this.chart, ctx = chartInstance.ctx; ctx.font = Chart.helpers.fontString(Chart.defaults.global.defaultFontSize-1, Chart.defaults.global.defaultFontStyle, Chart.defaults.global.defaultFontFamily); ctx.textAlign = 'center'; ctx.textBaseline = 'bottom'; this.data.datasets.forEach(function (dataset, i) { var meta = chartInstance.controller.getDatasetMeta(i); meta.data.forEach(function (bar, index) { var data = dataset.data[index]; ctx.fillText(data, bar._model.x, bar._model.y); });});}},
	legend: {display: false},title: {display: true,text: 'Número de casos por data de coleta'},maintainAspectRatio: false,scales: {yAxes: [{ticks: {beginAtZero: true}}]}}});
	*/
	var ctxB = document.getElementById("GraphTotalCasos").getContext('2d');
	var myBarChart = new Chart(ctxB, {type: 'bar',data: {	labels: ["29/02", "01/03", "09/03", "11/03", "12/03", "13/03", "14/03", "15/03", "16/03", "17/03", "18/03", "19/03", "20/03", "21/03", "22/03", "23/03", "24/03", "25/03", "26/03", "27/03"],	datasets: [{	label: 'total',	data: [1,2,5,7,11,16,26,33,62,82,108,132,161,168,180,197,217,226,232,239],	backgroundColor: 'rgba(166, 58, 58, 0.8)',	borderColor: 'rgba(0,0,0,1)',	borderWidth: 1	}]},options: {
	"hover": {"animationDuration": 0},"animation": { "duration": 1,"onComplete": function () { var chartInstance = this.chart, ctx = chartInstance.ctx; ctx.font = Chart.helpers.fontString(Chart.defaults.global.defaultFontSize-1, Chart.defaults.global.defaultFontStyle, Chart.defaults.global.defaultFontFamily); ctx.textAlign = 'center'; ctx.textBaseline = 'bottom'; this.data.datasets.forEach(function (dataset, i) { var meta = chartInstance.controller.getDatasetMeta(i); meta.data.forEach(function (bar, index) { var data = dataset.data[index]; ctx.fillText(data, bar._model.x, bar._model.y); });});}},
	legend: {display: false},title: {display: true,text: 'Número total de confirmados por data de coleta'},maintainAspectRatio: false,scales: {yAxes: [{ticks: {beginAtZero: true}}]}}});
	
	var ctxP = document.getElementById("GraphSexo").getContext('2d');var myPieChart = new Chart(ctxP, {type: 'pie',data: {labels: ["Masculino", "Feminino"],datasets: [{data: [147, 92],backgroundColor: ["#538B21", "#F73A30", "#ffffff",],hoverBackgroundColor: ["#79D826", "#FB847D", "#ffffff"]}]},options: {maintainAspectRatio: false,title: {display: true,text: 'Sexo'},responsive: true}});
	
	var ctx = document.getElementById("GraphIdade").getContext('2d');var myChart = new Chart(ctx, {type: 'horizontalBar',data: {labels: ["05-09", "10-14", "15-19", "20-29", "30-39", "40-49", "50-59", "60-69", "70-79", "80 e mais"],datasets: [{data: [1,1,5,30,50,38,47,40,18,9],backgroundColor: 'rgba(78, 170, 28, 0.8)',borderColor: 'rgba(0,0,0,1)',borderWidth: 1}]},options: {
	"hover": {"animationDuration": 0},"animation": { "duration": 1,"onComplete": function () { var chartInstance = this.chart, ctx = chartInstance.ctx; ctx.font = Chart.helpers.fontString(Chart.defaults.global.defaultFontSize-1, Chart.defaults.global.defaultFontStyle, Chart.defaults.global.defaultFontFamily); ctx.textAlign = 'center'; ctx.textBaseline = 'bottom'; this.data.datasets.forEach(function (dataset, i) { var meta = chartInstance.controller.getDatasetMeta(i); meta.data.forEach(function (bar, index) { var data = dataset.data[index]; ctx.fillText(data, bar._model.x + 10, bar._model.y + 7); });});}},
	maintainAspectRatio: false,legend: {display: false},title: {display: true,text: 'Faixa Etária'},scales: {yAxes: [{ticks: {beginAtZero: true}}]}}});
	
	function beforePrintHandler () {
		for (var id in Chart.instances) {
			Chart.instances[id].resize();
		}
	}
	
	var markers = [
			{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Bagé", lat: -31.32965, lng: -54.09992},{"name":" Bagé", lat: -31.32965, lng: -54.09992},{"name":" Bagé", lat: -31.32965, lng: -54.09992},{"name":" Bagé", lat: -31.32965, lng: -54.09992},{"name":" Bagé", lat: -31.32965, lng: -54.09992},{"name":" Bagé", lat: -31.32965, lng: -54.09992},{"name":" Bagé", lat: -31.32965, lng: -54.09992},{"name":" Bagé", lat: -31.32965, lng: -54.09992},{"name":" Bagé", lat: -31.32965, lng: -54.09992},{"name":" Canoas", lat: -29.912758, lng: -51.185681},{"name":" Canoas", lat: -29.912758, lng: -51.185681},{"name":" Canoas", lat: -29.912758, lng: -51.185681},{"name":" Canoas", lat: -29.912758, lng: -51.185681},{"name":" Canoas", lat: -29.912758, lng: -51.185681},{"name":" Canoas", lat: -29.912758, lng: -51.185681},{"name":" Canoas", lat: -29.912758, lng: -51.185681},{"name":" Canoas", lat: -29.912758, lng: -51.185681},{"name":" Torres", lat: -29.333438, lng: -49.733289},{"name":" Torres", lat: -29.333438, lng: -49.733289},{"name":" Torres", lat: -29.333438, lng: -49.733289},{"name":" Torres", lat: -29.333438, lng: -49.733289},{"name":" Torres", lat: -29.333438, lng: -49.733289},{"name":" Torres", lat: -29.333438, lng: -49.733289},{"name":" Torres", lat: -29.333438, lng: -49.733289},{"name":" Caxias do Sul", lat: -29.162905, lng: -51.179161},{"name":" Caxias do Sul", lat: -29.162905, lng: -51.179161},{"name":" Caxias do Sul", lat: -29.162905, lng: -51.179161},{"name":" Caxias do Sul", lat: -29.162905, lng: -51.179161},{"name":" Caxias do Sul", lat: -29.162905, lng: -51.179161},{"name":" Caxias do Sul", lat: -29.162905, lng: -51.179161},{"name":" Lajeado", lat: -29.459086, lng: -51.964427},{"name":" Lajeado", lat: -29.459086, lng: -51.964427},{"name":" Lajeado", lat: -29.459086, lng: -51.964427},{"name":" Lajeado", lat: -29.459086, lng: -51.964427},{"name":" Lajeado", lat: -29.459086, lng: -51.964427},{"name":" Sant'Ana do Livramento", lat: -30.7385244041684, lng: -55.5614082776956},{"name":" Sant'Ana do Livramento", lat: -30.7385244041684, lng: -55.5614082776956},{"name":" Sant'Ana do Livramento", lat: -30.7385244041684, lng: -55.5614082776956},{"name":" Sant'Ana do Livramento", lat: -30.7385244041684, lng: -55.5614082776956},{"name":" Serafina Corrêa", lat: -28.712621, lng: -51.93525},{"name":" Serafina Corrêa", lat: -28.712621, lng: -51.93525},{"name":" Serafina Corrêa", lat: -28.712621, lng: -51.93525},{"name":" Ivoti", lat: -29.599463, lng: -51.153326},{"name":" Ivoti", lat: -29.599463, lng: -51.153326},{"name":" Ivoti", lat: -29.599463, lng: -51.153326},{"name":" São Leopoldo", lat: -29.754494, lng: -51.149773},{"name":" São Leopoldo", lat: -29.754494, lng: -51.149773},{"name":" São Leopoldo", lat: -29.754494, lng: -51.149773},{"name":" Erechim", lat: -27.63638, lng: -52.26969},{"name":" Erechim", lat: -27.63638, lng: -52.26969},{"name":" Erechim", lat: -27.63638, lng: -52.26969},{"name":" Capão da Canoa", lat: -29.764242, lng: -50.028243},{"name":" Capão da Canoa", lat: -29.764242, lng: -50.028243},{"name":" Pelotas", lat: -31.764898, lng: -52.337058},{"name":" Pelotas", lat: -31.764898, lng: -52.337058},{"name":" Rio Grande", lat: -32.034875, lng: -52.10705},{"name":" Rio Grande", lat: -32.034875, lng: -52.10705},{"name":" Campo Bom", lat: -29.674694, lng: -51.060601},{"name":" Campo Bom", lat: -29.674694, lng: -51.060601},{"name":" Santa Maria", lat: -29.686817, lng: -53.814946},{"name":" Santa Maria", lat: -29.686817, lng: -53.814946},{"name":" Anta Gorda", lat: -28.969772, lng: -52.010191},{"name":" Anta Gorda", lat: -28.969772, lng: -52.010191},{"name":" Estância Velha", lat: -29.65354, lng: -51.184339},{"name":" Estância Velha", lat: -29.65354, lng: -51.184339},{"name":" Bento Gonçalves", lat: -29.166212, lng: -51.516476},{"name":" Bento Gonçalves", lat: -29.166212, lng: -51.516476},{"name":" Alvorada", lat: -29.9914, lng: -51.080857},{"name":" Alvorada", lat: -29.9914, lng: -51.080857},{"name":" Paraí", lat: -28.596384, lng: -51.789602},{"name":" Osório", lat: -29.888089, lng: -50.266713},{"name":" Dom Pedrito", lat: -30.97559, lng: -54.66936},{"name":" Marau", lat: -28.449754, lng: -52.198625},{"name":" Piratini", lat: -31.447291, lng: -53.097296},{"name":" Cachoeira do Sul", lat: -30.032988, lng: -52.892756},{"name":" Novo Hamburgo", lat: -29.687548, lng: -51.132828},{"name":" Canguçu", lat: -31.39598, lng: -52.678254},{"name":" Santa Rosa", lat: -27.8702, lng: -54.479629},{"name":" Nova Palma", lat: -29.470981, lng: -53.468924},{"name":" Rolante", lat: -29.64624, lng: -50.581937},{"name":" Garibaldi", lat: -29.258979, lng: -51.535179},{"name":" Passo Fundo", lat: -28.257564, lng: -52.409112},{"name":" Sapiranga", lat: -29.634885, lng: -51.006446},{"name":" Eldorado do Sul", lat: -30.084676, lng: -51.618702},{"name":" Viamão", lat: -30.081899, lng: -51.019435},{"name":" Dois Irmãos", lat: -29.58356, lng: -51.089776},{"name":" Charqueadas", lat: -29.962478, lng: -51.628872},{"name":" Taquara", lat: -29.650471, lng: -50.775278},{"name":" Carlos Barbosa", lat: -29.296949, lng: -51.502847},{"name":" Estrela", lat: -29.501611, lng: -51.960548},{"name":" Cruzeiro do Sul", lat: -29.514835, lng: -51.992778},{"name":" Cerro Grande do Sul", lat: -30.590472, lng: -51.741846},{"name":" Santiago", lat: -29.189675, lng: -54.866624},{"name":" Farroupilha", lat: -29.222689, lng: -51.341853},{"name":" Santo Antônio da Patrulha", lat: -29.826768, lng: -50.517489},{"name":" Gravataí", lat: -29.941319, lng: -50.986891},];	
	
	var markersObitos = [
			{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},{"name":" Porto Alegre", lat: -30.031771, lng: -51.206533},];	
			
	var map = L.map( 'map', { center: [-30.5, -53.0], minZoom: 6, maxZoom: 10,  zoom: 6});
	L.tileLayer( 'http://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
	 attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> | SES-RS/CEVS/DGTI | Atualizado em 29/03/2020 20:18:12',
	 subdomains: ['a','b','c']
	}).addTo( map );
	
	
	var myIcon = L.icon({
	  iconUrl: 'http://ti.saude.rs.gov.br/covid19/leaflet/images/marker.png',
	  iconRetinaUrl:  'http://ti.saude.rs.gov.br/covid19/leaflet/images/marker.png',
	  iconSize: [25, 41],
	  iconAnchor: [9, 21],
	  popupAnchor: [0, -14]
	});
	var markerClusters = L.markerClusterGroup({maxClusterRadius:20});
	for ( var i = 0; i < markers.length; ++i )
	{
	  var popup = markers[i].name;
	  var m = L.marker( [markers[i].lat, markers[i].lng], {icon: myIcon} ).bindPopup( popup );
	  markerClusters.addLayer( m );
	}
	map.addLayer( markerClusters );
	
	
	var mapObitos = L.map( 'mapObitos', { center: [-30.5, -53.0], minZoom: 6, maxZoom: 10,  zoom: 6});
	L.tileLayer( 'http://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
	 attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> | SES-RS/CEVS/DGTI | Atualizado em 29/03/2020 20:18:12',
	 subdomains: ['a','b','c']
	}).addTo( mapObitos );
	
	var markerClustersObitos = L.markerClusterGroup({maxClusterRadius:20});
	for ( var i = 0; i < markersObitos.length; ++i )
	{
	  var popup = markersObitos[i].name;
	  var m = L.marker( [markersObitos[i].lat, markersObitos[i].lng], {icon: myIcon} ).bindPopup( popup );
	  markerClustersObitos.addLayer( m );
	}
	mapObitos.addLayer( markerClustersObitos );
	
	function trocamapa(x, map, mapObitos)
	{
		if(x==1) {
			$("#map").css("display","");
			$("#mapObitos").css("display","none");
			$("#btnConf").css('background-color',"#ffffff");
			$("#btnObitos").css('background-color',"#dfdfdf");
			setTimeout(function(){ map.invalidateSize()}, 400);
		} else {
			$("#map").css("display","none");
			$("#mapObitos").css("display","");
			setTimeout(function(){ mapObitos.invalidateSize()}, 400);
			$("#btnConf").css('background-color',"#dfdfdf");
			$("#btnObitos").css('background-color',"#ffffff");
		}
	}
	
	</script>
	
	  
	
	</body></html>`
}

func TestReplaceValue(t *testing.T) {
	type args struct {
		splitCity []string
		splitVals []string
	}
	tt := struct {
		name       string
		args       args
		wantCities []string
		wantValues []string
	}{
		name: "Test Replace string to cities and values",
		args: args{
			splitCity: []string{"[porto alegre", "caxias]", "d"},
			splitVals: []string{"[1", "2]"},
		},
		wantCities: []string{"porto alegre", "caxias"},
		wantValues: []string{"1", "2"},
	}
	t.Run(tt.name, func(t *testing.T) {
		gotCities, gotValues := ReplaceValue(tt.args.splitCity, tt.args.splitVals)
		if !reflect.DeepEqual(gotCities, tt.wantCities) {
			t.Errorf("ReplaceValue() got = %v, want %v", gotCities, tt.wantCities)
		}
		if !reflect.DeepEqual(gotValues, tt.wantValues) {
			t.Errorf("ReplaceValue() got1 = %v, want %v", gotValues, tt.wantValues)
		}
	})
}

func TestGetCityValue(t *testing.T) {
	mockBody()
	tt := struct {
		name   string
		body   string
		cities []string
		values []string
	}{
		name:   "Test GetCityValue in body",
		body:   body,
		cities: []string{"[\" Porto Alegre\"", " \" Bagé\"", " \" Canoas\"]", "d"},
		values: []string{"[143", "9", "8]"},
	}
	t.Run(tt.name, func(t *testing.T) {
		cities, values := GetCityValue(tt.body)
		if !reflect.DeepEqual(cities, tt.cities) {
			t.Errorf("GetCityValue() got = %v, want %v", cities, tt.cities)
		}
		if !reflect.DeepEqual(values, tt.values) {
			t.Errorf("GetCityValue() got1 = %v, want %v", values, tt.values)
		}
	})
}

func TestSetCitiesValues(t *testing.T) {
	type args struct {
		splitCity []string
		splitVals []string
	}
	tt := struct {
		name         string
		args         args
		citiesValues ListCityValue
	}{
		name: "Test Set Cities and Values in struct",
		args: args{
			splitCity: []string{"Porto Alegre", "Bagé", "Canoas"},
			splitVals: []string{"143", "9", "8"},
		},
		citiesValues: ListCityValue{
			data: []CityValue{
				{
					City:  "Porto Alegre",
					Value: 143,
				}, {
					City:  "Bagé",
					Value: 9,
				}, {
					City:  "Canoas",
					Value: 8,
				},
			},
			total: 160,
		},
	}
	t.Run(tt.name, func(t *testing.T) {
		if got := SetCitiesValues(tt.args.splitCity, tt.args.splitVals); !reflect.DeepEqual(got, tt.citiesValues) {
			t.Errorf("SetCitiesValues() = %v, want %v", got, tt.citiesValues)
		}
	})
}
