package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"html/template"
	"hash/crc32"
	"os"
	"strings"
	"time"
)

type Page struct {
    Title string
    Ip  string
    Mac string
    BgColor int
    FgColor int
    Request *http.Request
}


var tmplGraph = `<html>
  <head>
    <title>{{.Title}}</title>
    <style>

#r-two-d-two-space {
background:#{{printf "%x" .BgColor}};
width:600px;
padding-top:60px;
}

#r-two-d-two {
width:225px;	
margin:auto;
}

#r-two-d-two-head {
height:140px;
background: #ccc; 
background: -moz-linear-gradient(left, #606060 0%, #f9f9f9 50%, #6d6d6d 100%); 
background: -webkit-gradient(linear, left top, right top, color-stop(0%,#606060), color-stop(50%,#f9f9f9), color-stop(100%,#6d6d6d)); 
background: -webkit-linear-gradient(left, #606060 0%,#f9f9f9 50%,#6d6d6d 100%); 
background: -o-linear-gradient(left, #606060 0%,#f9f9f9 50%,#6d6d6d 100%); 
background: -ms-linear-gradient(left, #606060 0%,#f9f9f9 50%,#6d6d6d 100%); 
background: linear-gradient(to right, #606060 0%,#f9f9f9 50%,#6d6d6d 100%); 
width:229px;
padding-top:5px;
padding-bottom:5px;
margin-left:-35px;
margin-bottom:-3px;
-webkit-border-top-right-radius: 95px;
-webkit-border-top-left-radius: 95px;
-moz-border-radius-topright: 95px;
-moz-border-radius-topleft: 95px;
border-top-right-radius: 95px;
border-top-left-radius: 95px;
-moz-transform: rotate(-10deg);
-webkit-transform: rotate(-10deg);
-o-transform: rotate(-10deg);
-ms-transform: rotate(-10deg);
transform: rotate(-10deg);
position:relative;
z-index:99;
-webkit-box-shadow:inset 0 22px 22px rgba(0, 0, 0, .5);
-moz-box-shadow:inset 0 22px 22px rgba(0, 0, 0, .5);
box-shadow:inset 0 22px 22px rgba(0, 0, 0, .5);
}

#r-two-d-two-body {
height:200px;
background: #bfbfbf; 
background: -moz-linear-gradient(left, #bfbfbf 0%, #dbdbdb 8%, #ffffff 35%, #ffffff 61%, #b7b7b7 100%); 
background: -webkit-gradient(linear, left top, right top, color-stop(0%,#bfbfbf), color-stop(8%,#dbdbdb), color-stop(35%,#ffffff), color-stop(61%,#ffffff), color-stop(100%,#b7b7b7)); 
background: -webkit-linear-gradient(left, #bfbfbf 0%,#dbdbdb 8%,#ffffff 35%,#ffffff 61%,#b7b7b7 100%); 
background: -o-linear-gradient(left, #bfbfbf 0%,#dbdbdb 8%,#ffffff 35%,#ffffff 61%,#b7b7b7 100%); 
background: -ms-linear-gradient(left, #bfbfbf 0%,#dbdbdb 8%,#ffffff 35%,#ffffff 61%,#b7b7b7 100%); 
background: linear-gradient(to right, #bfbfbf 0%,#dbdbdb 8%,#ffffff 35%,#ffffff 61%,#b7b7b7 100%); 
width:228px;
margin-left:-4px;
padding-top:5px;
position:relative;
z-index:99;
-moz-transform: rotate(-10deg);
-webkit-transform: rotate(-10deg);
-o-transform: rotate(-10deg);
-ms-transform: rotate(-10deg);
transform: rotate(-10deg);
}

#r-two-d-two-trim {
width: 180px;
height: 0;
margin-top:-3px;
margin-left:14px;
border-left: 25px solid transparent;
border-top: 20px solid #fff;
border-right: 25px solid transparent;
-moz-transform: rotate(-10deg);
-webkit-transform: rotate(-10deg);
-o-transform: rotate(-10deg);
-ms-transform: rotate(-10deg);
transform: rotate(-10deg);
position:relative;
z-index:99;
}

#r-two-d-two-leg-back {
height:30px;
background:#ccc;
border-left:5px solid #777;
width:60px;
padding-top:10px;
position:relative;
-webkit-border-bottom-right-radius: 90px;
-webkit-border-bottom-left-radius: 90px;
-moz-border-radius-bottomright: 90px;
-moz-border-radius-bottomleft: 90px;
border-bottom-right-radius: 90px;
border-bottom-left-radius: 90px;
margin-left:55px;
}

#r-two-d-two-innerleg-back {
height:20px;
background:#999;
width:40px;
border-bottom:1px solid #999;
border-left:1px solid #999;
border-right:1px solid #999;
border-top:1px solid #999;
position:relative;
-webkit-border-top-right-radius: 95px;
-webkit-border-top-left-radius: 95px;
-moz-border-radius-topright: 95px;
-moz-border-radius-topleft: 95px;
border-top-right-radius: 95px;
border-top-left-radius: 95px;
margin-left:auto;
margin-right:auto;
-webkit-box-shadow:inset 0 2px 2px rgba(0, 0, 0, .5);
-moz-box-shadow:inset 0 2px 2px rgba(0, 0, 0, .5);
box-shadow:inset 0 2px 2px rgba(0, 0, 0, .5);
}

#r-two-d-two-leg-front {
height:230px;
width:70px;
position:relative;
-webkit-border-top-right-radius: 95px;
-webkit-border-top-left-radius: 95px;
-moz-border-radius-topright: 95px;
-moz-border-radius-topleft: 95px;
border-top-right-radius: 95px;
border-top-left-radius: 95px;
-webkit-border-bottom-right-radius: 20px;
-webkit-border-bottom-left-radius: 20px;
-moz-border-radius-bottomright: 20px;
-moz-border-radius-bottomleft: 20px;
border-bottom-right-radius: 20px;
border-bottom-left-radius: 20px;
margin-left:auto;
margin-right:auto;
position:relative;
z-index:99;
}

#r-two-d-two-leg-top-square {
height:18px;
width:18px;
border:2px solid #444;
-webkit-border-radius: 95px;
-moz-border-radius: 95px;
border-radius: 95px;
margin-left:auto;
margin-right:auto;
margin-top:10px;
-moz-box-shadow:0 0 3px rgba(255, 255, 255, .5);
-webkit-box-shadow:0 0 3px rgba(255, 255, 255, .5);
box-shadow: 0 0 3px rgba(255, 255, 255, .5);
}

#r-two-d-two-innerleg-top {
height:40px;
background: #606060; 
background: -moz-linear-gradient(left, #606060 0%, #f9f9f9 50%, #6d6d6d 100%); 
background: -webkit-gradient(linear, left top, right top, color-stop(0%,#606060), color-stop(50%,#f9f9f9), color-stop(100%,#6d6d6d)); 
background: -webkit-linear-gradient(left, #606060 0%,#f9f9f9 50%,#6d6d6d 100%); 
background: -o-linear-gradient(left, #606060 0%,#f9f9f9 50%,#6d6d6d 100%); 
background: -ms-linear-gradient(left, #606060 0%,#f9f9f9 50%,#6d6d6d 100%); 
background: linear-gradient(to right, #606060 0%,#f9f9f9 50%,#6d6d6d 100%);
width:40px;
margin-bottom:-10px;
position:relative;
z-index:999;
border:2px solid #eee;
-webkit-border-radius: 95px;
-moz-border-radius: 95px;
border-radius: 95px;
margin-left:auto;
margin-right:auto;
-webkit-box-shadow:inset 0 2px 2px rgba(0, 0, 0, .5);
-moz-box-shadow:inset 0 2px 2px rgba(0, 0, 0, .5);
box-shadow:inset 0 2px 2px rgba(0, 0, 0, .5);
}

#r-two-d-two-innerleg-front {
height:140px;
background:#eee;
width:44px;
margin-left:auto;
margin-right:auto;
-webkit-box-shadow: 0 1px 3px rgba(0, 0, 0, .6);
-moz-box-shadow: 0 1px 3px rgba(0, 0, 0, .6);
box-shadow:0 1px 3px rgba(0, 0, 0, .6);
position:relative;
z-index:9;
}

#r-two-d-two-leg-ankle {
height:42px;
background:#eee;
width:64px;
padding-top:10px;
border-left:5px solid #ccc;
margin-left:auto;
margin-right:auto;
-webkit-border-bottom-right-radius: 25px;
-webkit-border-bottom-left-radius:25px;
-moz-border-radius-bottomright: 25px;
-moz-border-radius-bottomleft: 25px;
border-bottom-right-radius: 25px;
border-bottom-left-radius:25px;
-webkit-box-shadow: 0 1px 3px rgba(0, 0, 0, .6);
-moz-box-shadow: 0 1px 3px rgba(0, 0, 0, .6);
box-shadow:0 1px 3px rgba(0, 0, 0, .6);
position:relative;
z-index:99;
}

#r-two-d-two-leg-anklein {
height:20px;
background:#{{printf "%x" .FgColor}};
width:34px;
border-top:1px solid #eee;
border-left:1px solid #eee;
margin-left:auto;
margin-right:auto;
-webkit-box-shadow:inset  0 4px 8px rgba(0, 0, 0, .6);
-moz-box-shadow:inset 0 4px 8px rgba(0, 0, 0, .6);
box-shadow:inset 0 4px 8px rgba(0, 0, 0, .6);
position:relative;
z-index:99;
}

#r-two-d-two-leg-anklebar {
height:7px;
background: #bfbfbf; 
background: -moz-linear-gradient(top, #bfbfbf 0%, #dbdbdb 8%, #ffffff 35%, #ffffff 61%, #b7b7b7 100%); 
background: -webkit-gradient(linear, left top, left bottom, color-stop(0%,#bfbfbf), color-stop(8%,#dbdbdb), color-stop(35%,#ffffff), color-stop(61%,#ffffff), color-stop(100%,#b7b7b7)); 
background: -webkit-linear-gradient(top, #bfbfbf 0%,#dbdbdb 8%,#ffffff 35%,#ffffff 61%,#b7b7b7 100%); 
background: -o-linear-gradient(top, #bfbfbf 0%,#dbdbdb 8%,#ffffff 35%,#ffffff 61%,#b7b7b7 100%); 
background: -ms-linear-gradient(top, #bfbfbf 0%,#dbdbdb 8%,#ffffff 35%,#ffffff 61%,#b7b7b7 100%); 
background: linear-gradient(to bottom, #bfbfbf 0%,#dbdbdb 8%,#ffffff 35%,#ffffff 61%,#b7b7b7 100%); 
margin-top:10px;
margin-bottom:-15px;
width:63px;
border-bottom:1px solid #999;
border-left:1px solid #999;
border-right:1px solid #999;
border-top:1px solid #999;
position:relative;
z-index:99;
-webkit-border-radius: 95px;
-moz-border-radius: 95px;
border-radius: 95px;
margin-left:auto;
margin-right:auto;
-webkit-box-shadow:0 2px 2px rgba(0, 0, 0, .5);
-moz-box-shadow: 0 2px 2px rgba(0, 0, 0, .5);
box-shadow: 0 2px 2px rgba(0, 0, 0, .5);
}

#r-two-d-two-innerleg-front-pipe {
height:115px;
background:#{{printf "%x" .FgColor}};
margin-top:3px;
width:30px;
border-bottom:1px solid #999;
border-left:1px solid #999;
border-right:1px solid #999;
border-top:1px solid #999;
position:relative;
z-index:9999;
-webkit-border-top-right-radius: 95px;
-webkit-border-top-left-radius: 95px;
-moz-border-radius-topright: 95px;
-moz-border-radius-topleft: 95px;
border-top-right-radius: 95px;
border-top-left-radius: 95px;
margin-left:auto;
margin-right:auto;
-webkit-box-shadow:inset 0 5px 6px rgba(0, 0, 0, .5);
-moz-box-shadow:inset 0 5px 6px rgba(0, 0, 0, .5);
box-shadow:inset 0 5px 6px rgba(0, 0, 0, .5);
}

#r-two-d-two-innerleg-front-pipein {
height:83px;
background: #606060; 
background: -moz-linear-gradient(left, #606060 0%, #f9f9f9 50%, #6d6d6d 100%); 
background: -webkit-gradient(linear, left top, right top, color-stop(0%,#606060), color-stop(50%,#f9f9f9), color-stop(100%,#6d6d6d)); 
background: -webkit-linear-gradient(left, #606060 0%,#f9f9f9 50%,#6d6d6d 100%); 
background: -o-linear-gradient(left, #606060 0%,#f9f9f9 50%,#6d6d6d 100%); 
background: -ms-linear-gradient(left, #606060 0%,#f9f9f9 50%,#6d6d6d 100%); 
background: linear-gradient(to right, #606060 0%,#f9f9f9 50%,#6d6d6d 100%); margin-top:30px;
width:7px;
border-bottom:1px solid #999;
border-left:1px solid #222;
border-right:1px solid #eee;
border-top:1px solid #999;
position:relative;
-webkit-border-top-right-radius: 95px;
-webkit-border-top-left-radius: 95px;
-moz-border-radius-topright: 95px;
-moz-border-radius-topleft: 95px;
border-top-right-radius: 95px;
border-top-left-radius: 95px;
margin-left:auto;
margin-right:auto;
-webkit-box-shadow:inset 0 2px 2px rgba(0, 0, 0, .5);
-moz-box-shadow:inset 0 2px 2px rgba(0, 0, 0, .5);
box-shadow:inset 0 2px 2px rgba(0, 0, 0, .5);
}

#r-two-d-two-foot-front {
width: 80px;
height: 0;
margin-bottom:3px;
margin-left:auto;
margin-right:auto;
border-left: 50px solid transparent;
border-bottom: 50px solid #f1f1f1;
border-right: 50px solid transparent;
position:relative;
z-index:99;
}

#r-two-d-two-foot-back {
width: 80px;
height: 0;
margin-left:18px;
border-left: 40px solid transparent;
border-bottom: 40px solid #ccc;
border-right: 40px solid transparent;
position:relative;
z-index:9;
}

#r-two-d-two-lowerfoot-round {
height:30px;
width:80px;
margin-top:-35px;
background:#ccc;
-webkit-border-top-right-radius: 95px;
-webkit-border-top-left-radius: 95px;
-moz-border-radius-topright: 95px;
-moz-border-radius-topleft: 95px;
border-top-right-radius: 95px;
border-top-left-radius: 95px;
position:relative;
margin-left:auto;
margin-right:auto;
z-index:99;
-webkit-box-shadow:inset 0 2px 6px rgba(0, 0, 0, .5);
-moz-box-shadow:inset 0 2px 6px rgba(0, 0, 0, .5);
box-shadow:inset 0 2px 6px rgba(0, 0, 0, .5);
}

#r-two-d-two-lowerfoot-front {
height:20px;
background:#ccc;
width:180px;
position:relative;
margin-left:auto;
margin-right:auto;
-webkit-box-shadow: 0 1px 3px rgba(0, 0, 0, .6);
-moz-box-shadow: 0 1px 3px rgba(0, 0, 0, .6);
box-shadow:0 1px 3px rgba(0, 0, 0, .6);
}

#r-two-d-two-lowerfoot-back {
height:20px;
background:#999;
width:160px;
margin-left:18px;
position:relative;
-webkit-box-shadow: 0 1px 3px rgba(0, 0, 0, .6);
-moz-box-shadow: 0 1px 3px rgba(0, 0, 0, .6);
box-shadow:0 1px 3px rgba(0, 0, 0, .6);
}

#r-two-d-two-head-backunit {
height:15px;
width:18px;
float:right;
margin-bottom:-8px;
background:#ccc;
margin-left:6px;
margin-right:74px;
border-top: 1px solid #ccc;
-webkit-box-shadow:inset 0 -12px 12px rgba(0, 0, 0, .5);
-moz-box-shadow:inset 0 -12px 12px rgba(0, 0, 0, .5);
box-shadow:inset 0 -12px 12px rgba(0, 0, 0, .5);
-moz-transform:rotate(22deg);
-webkit-transform:rotate(22deg);
-o-transform:rotate(22deg);
-ms-transform:rotate(22deg);
transform:rotate(22deg);
}

#r-two-d-two-head-projector-bluebox {
height:50px;
width:8px;
float:left;
margin-bottom:4px;
background:#{{printf "%x" .FgColor}};
margin-left:2px;
margin-right:4px;
margin-top:15px;
border-left:1px solid blue;
-webkit-box-shadow:inset 0 -12px 12px rgba(0, 0, 0, .8);
-moz-box-shadow:inset 0 -12px 12px rgba(0, 0, 0, .8);
box-shadow:inset 0 -12px 12px rgba(0, 0, 0, .8);
-moz-transform:rotate(17deg);
-webkit-transform:rotate(17deg);
-o-transform:rotate(17deg);
-ms-transform:rotate(17deg);
transform:rotate(17deg);
}

#r-two-d-two-head-projector-bulb {
height:30px;
width:12px;
float:left;
margin-bottom:4px;
background: #bfbfbf; 
background: -moz-linear-gradient(top, #bfbfbf 0%, #dbdbdb 8%, #ffffff 35%, #ffffff 61%, #b7b7b7 100%); 
background: -webkit-gradient(linear, left top, left bottom, color-stop(0%,#bfbfbf), color-stop(8%,#dbdbdb), color-stop(35%,#ffffff), color-stop(61%,#ffffff), color-stop(100%,#b7b7b7)); 
background: -webkit-linear-gradient(top, #bfbfbf 0%,#dbdbdb 8%,#ffffff 35%,#ffffff 61%,#b7b7b7 100%); 
background: -o-linear-gradient(top, #bfbfbf 0%,#dbdbdb 8%,#ffffff 35%,#ffffff 61%,#b7b7b7 100%); 
background: -ms-linear-gradient(top, #bfbfbf 0%,#dbdbdb 8%,#ffffff 35%,#ffffff 61%,#b7b7b7 100%); 
background: linear-gradient(to bottom, #bfbfbf 0%,#dbdbdb 8%,#ffffff 35%,#ffffff 61%,#b7b7b7 100%); 
margin-right:4px;
margin-top:64px;
margin-left:-25px;
border-left:1px dashed lightskyblue;
border-top:1px solid #ccc;
border-bottom:1px solid #444;
position:relative;
-webkit-box-shadow:inset 0 -12px 12px rgba(0, 0, 0, .5);
-moz-box-shadow:inset 0 -12px 12px rgba(0, 0, 0, .5);
box-shadow:inset 0 -12px 12px rgba(0, 0, 0, .5);
}

.r-two-d-two-head-sensor {
height:30px;
width:19px;
float:left;
margin-bottom:4px;
background:#{{printf "%x" .FgColor}};
margin-top:60px;
margin-left:9px;
border-left:1px solid darkblue;
border-top:1px solid darkblue;
border-bottom:1px solid #ccc;
border-right:1px solid #ccc;
-webkit-box-shadow:inset 0 12px 12px rgba(0, 0, 0, .5);
-moz-box-shadow:inset 0 12px 12px rgba(0, 0, 0, .5);
box-shadow:inset 0 12px 12px rgba(0, 0, 0, .5);
}

#r-two-d-two-head-bigsensor {
height:30px;
width:32px;
float:left;
background:#{{printf "%x" .FgColor}};
margin-right:4px;
margin-top:60px;
margin-left:9px;
border-left:1px solid darkblue;
border-top:1px solid darkblue;
border-bottom:1px solid #fff;
border-right:1px solid #fff;
-webkit-box-shadow:inset 0 12px 12px rgba(0, 0, 0, .5);
-moz-box-shadow:inset 0 12px 12px rgba(0, 0, 0, .5);
box-shadow:inset 0 12px 12px rgba(0, 0, 0, .5);
}

#r-two-d-two-head-sensor-back {
height:30px;
width:22px;
float:left;
background:#{{printf "%x" .FgColor}};
margin-right:4px;
margin-top:60px;
margin-left:9px;
border-left:1px solid darkblue;
border-top:1px solid darkblue;
border-bottom:1px solid #fff;
border-right:1px solid #fff;
-webkit-box-shadow:inset 0 12px 12px rgba(0, 0, 0, .5);
-moz-box-shadow:inset 0 12px 12px rgba(0, 0, 0, .5);
box-shadow:inset 0 12px 12px rgba(0, 0, 0, .5);
}

#r-two-d-two-head-sensor-backhollow {
height:25px;
width:17px;
float:left;
margin-right:4px;
margin-top:60px;
margin-left:9px;
border-left:1px solid #888;
border-top:1px solid #888;
border-bottom:1px solid #ccc;
border-right:1px solid #ccc;
}

#r-two-d-two-head-sheen {
width:120px;
height:140px;
float:right;
margin-top:-10px;
margin-left:-450px;
margin-right:-32px;
margin-bottom:-113px;
border-right:70px solid #000;
position:relative;
z-index:9999;
opacity:.2;
-webkit-border-top-left-radius:0;
-webkit-border-top-right-radius:550px;
-moz-border-radius-topleft:0;
-moz-border-radius-topright:550px;
border-top-left-radius:0;
border-top-right-radius:550px;
-moz-transform:rotate(45deg);
-webkit-transform:rotate(45deg);
-o-transform:rotate(45deg);
-ms-transform:rotate (45deg);
transform:rotate(45deg);
}

#r-two-d-two-body-sheen {
width:80px;
height:180px;
float:right;
margin-top:-130px;
margin-left:50px;
margin-right:43px;
border-right:60px solid #999;
opacity:.2;
-webkit-border-bottom-left-radius:0;
-webkit-border-bottom-right-radius:550px;
-moz-border-radius-bottomleft:0;
-moz-border-radius-bottomright:550px;
border-bottom-left-radius:0;
border-bottom-right-radius:550px;
-moz-transform:rotate(45deg);
-webkit-transform:rotate(45deg);
-o-transform:rotate(45deg);
-ms-transform:rotate (45deg);
transform:rotate(45deg);
}

#r-two-d-two-head-sensorband {
height:8px;
width:228px;
clear:both;
background:#{{printf "%x" .FgColor}};
margin-bottom:20px;
margin-right:4px;
border-left:1px solid darkblue;
border-bottom:1px solid #eee;
border-top:1px solid #555;
-webkit-box-shadow:inset 0 12px 12px rgba(0, 0, 0, .5);
-moz-box-shadow:inset 0 12px 12px rgba(0, 0, 0, .5);
box-shadow:inset 0 12px 12px rgba(0, 0, 0, .5);
}

#r-two-d-two-body-back-one {
height:136px;
width:9px;
float: right;
margin-right:4px;
border-left:1px solid #999;
border-top:1px solid #999;
border-bottom:1px solid #eee;
border-right:1px solid #eee;
-webkit-box-shadow:inset 0 12px 12px rgba(0, 0, 0, .2);
-moz-box-shadow:inset 0 12px 12px rgba(0, 0, 0, .2);
box-shadow:inset 0 12px 12px rgba(0, 0, 0, .2);
}

#r-two-d-two-body-back-two {
height:136px;
width:19px;
float: right;
margin-right:4px;
border-left:1px solid #999;
border-top:1px solid #999;
border-bottom:1px solid #eee;
border-right:1px solid #eee;
-webkit-box-shadow:inset 0 12px 12px rgba(0, 0, 0, .2);
-moz-box-shadow:inset 0 12px 12px rgba(0, 0, 0, .2);
box-shadow:inset 0 12px 12px rgba(0, 0, 0, .2);
}

#r-two-d-two-body-back-three {
height:66px;
width:70px;
float: right;
margin-top:70px;
margin-bottom:10px;
margin-right:40px;
background: #606060; 
background: -moz-linear-gradient(left, #606060 0%, #f9f9f9 50%, #6d6d6d 100%); 
background: -webkit-gradient(linear, left top, right top, color-stop(0%,#606060), color-stop(50%,#f9f9f9), color-stop(100%,#6d6d6d)); 
background: -webkit-linear-gradient(left, #606060 0%,#f9f9f9 50%,#6d6d6d 100%); 
background: -o-linear-gradient(left, #606060 0%,#f9f9f9 50%,#6d6d6d 100%); 
background: -ms-linear-gradient(left, #606060 0%,#f9f9f9 50%,#6d6d6d 100%); 
background: linear-gradient(to right, #606060 0%,#f9f9f9 50%,#6d6d6d 100%); 
}

#r-two-d-two-body-back-four {
height:136px;
width:30px;
float: right;
margin-bottom:10px;
margin-right:14px;
border-left:1px solid #999;
border-top:1px solid #999;
border-bottom:1px solid #eee;
border-right:1px solid #eee;
-webkit-box-shadow:inset 0 12px 12px rgba(0, 0, 0, .2);
-moz-box-shadow:inset 0 12px 12px rgba(0, 0, 0, .2);
box-shadow:inset 0 12px 12px rgba(0, 0, 0, .2);
}

.r-two-d-two-body-back-threein {
height:63px;
width:8px;
margin-bottom:10px;
background:#{{printf "%x" .FgColor}};
float:left;
margin-left:7px;
margin-right:6px;
-webkit-box-shadow:inset 0 12px 12px rgba(0, 0, 0, .2);
-moz-box-shadow:inset 0 12px 12px rgba(0, 0, 0, .2);
box-shadow:inset 0 12px 12px rgba(0, 0, 0, .2);
-webkit-border-bottom-right-radius: 95px;
-webkit-border-bottom-left-radius: 95px;
-moz-border-radius-bottomright: 95px;
-moz-border-radius-bottomleft: 95px;
border-bottom-right-radius: 95px;
border-bottom-left-radius: 95px;
}

.r-two-d-two-body-sensor {
height:6px;
width:9px;
margin-bottom:10px;
background:#{{printf "%x" .FgColor}};
margin-right:4px;
border-left:1px solid #000;
border-top:1px solid #000;
border-bottom:1px solid #eee;
border-right:1px solid #eee;
-webkit-box-shadow:inset 0 12px 12px rgba(0, 0, 0, .2);
-moz-box-shadow:inset 0 12px 12px rgba(0, 0, 0, .2);
box-shadow:inset 0 12px 12px rgba(0, 0, 0, .2);
}

#r-two-d-two-body-bigspace{
width: 36px;
height: 37px;
position: relative;
margin-top:5px;
margin-left:10px;
opacity:.4;
}

#r-two-d-two-body-bigspace:before {
content: "";
position: absolute;
top: 0;
left: 0;    
border-bottom: 19px solid #999;
border-left: 8px solid transparent;
border-right: 8px solid transparent;
width: 32px;
height: 0;
}

#r-two-d-two-body-bigspace:after {
content: "";
position: absolute;
bottom: 0;
left: 0;    
border-top: 19px solid #999;
border-left: 8px solid transparent;
border-right: 8px solid transparent;
width: 32px;
height: 0;
}				

.r-two-d-two-body-lowersensor {
height:40px;
width:4px;
margin-bottom:10px;
background:#{{printf "%x" .FgColor}};
margin-right:4px;
border-left:1px solid blue;
border-top:1px solid #999;
border-bottom:1px solid #eee;
border-right:1px solid #eee;
-webkit-box-shadow:inset 0 12px 12px rgba(0, 0, 0, .5);
-moz-box-shadow:inset 0 12px 12px rgba(0, 0, 0, .5);
box-shadow:inset 0 12px 12px rgba(0, 0, 0, .5);
}

#r-two-d-two-body-lowersensor-midbox{
width:50px;
height:43px;
float:right;
padding-top: 3px;
margin-right: 5px;
background: #606060; 
background: -moz-linear-gradient(left, #606060 0%, #f9f9f9 50%, #6d6d6d 100%); 
background: -webkit-gradient(linear, left top, right top, color-stop(0%,#606060), color-stop(50%,#f9f9f9), color-stop(100%,#6d6d6d)); 
background: -webkit-linear-gradient(left, #606060 0%,#f9f9f9 50%,#6d6d6d 100%); 
background: -o-linear-gradient(left, #606060 0%,#f9f9f9 50%,#6d6d6d 100%); 
background: -ms-linear-gradient(left, #606060 0%,#f9f9f9 50%,#6d6d6d 100%); 
background: linear-gradient(to right, #606060 0%,#f9f9f9 50%,#6d6d6d 100%); 
-webkit-box-shadow:inset 0 22px 22px rgba(0, 0, 0, .3);
-moz-box-shadow:inset 0 22px 22px rgba(0, 0, 0, .3);
box-shadow:inset 0 22px 22px rgba(0, 0, 0, .3);
}

.r-two-d-two-body-lowersensor-mid {
height:30px;
width:8px;
margin-bottom:10px;
background:#{{printf "%x" .FgColor}};
float:right;
margin-left:2px;
border-left:1px solid #999;
border-top:1px solid #999;
border-bottom:1px solid #fff;
border-right:1px solid #fff;
}

#r-two-d-two-leg-top {
background:#eee;
width:90px;
height:110px;
margin-bottom:-90px;
border-left:5px solid #ccc;
-webkit-border-top-right-radius: 95px;
-webkit-border-top-left-radius: 95px;
-moz-border-radius-topright: 95px;
-moz-border-radius-topleft: 95px;
border-top-right-radius: 95px;
border-top-left-radius: 95px;
-webkit-border-bottom-right-radius: 18px;
-webkit-border-bottom-left-radius: 18px;
-moz-border-radius-bottomright: 18px;
-moz-border-radius-bottomleft: 18px;
border-bottom-right-radius: 18px;
border-bottom-left-radius: 18px;
margin-left:auto;
margin-right:auto;
position:relative;
z-index:99;
-webkit-box-shadow: 0 1px 3px rgba(0, 0, 0, .6);
-moz-box-shadow: 0 1px 3px rgba(0, 0, 0, .6);
box-shadow:0 1px 3px rgba(0, 0, 0, .6);
}

#r-two-d-two-leg-out {
height:370px;
position:relative;
z-index:999;
margin-top:-315px;
}

#r-two-d-two-head-top {
height:24px;
width:120px;
clear:both;
background:#{{printf "%x" .FgColor}};
margin-right:4px;
overflow:hidden;
margin-left:46px;
border-left:1px solid darkblue;
border-right:1px solid #fff;
-webkit-box-shadow:inset 0 12px 12px rgba(0, 0, 0, .5);
-moz-box-shadow:inset 0 12px 12px rgba(0, 0, 0, .5);
box-shadow:inset 0 12px 12px rgba(0, 0, 0, .5);
-webkit-border-top-right-radius: 95px;
-webkit-border-top-left-radius: 95px;
-moz-border-radius-topright: 95px;
-moz-border-radius-topleft: 95px;
border-top-right-radius: 95px;
border-top-left-radius: 95px;
}

#r-two-d-two-head-topin {
height:138px;
width:40px;
margin-top:-10px;
clear:both;
background:#{{printf "%x" .FgColor}};
margin-right:4px;
margin-left:36px;
border-left:3px solid #999;
border-right:3px solid #999;
-webkit-box-shadow:inset 0 12px 12px rgba(0, 0, 0, .5);
-moz-box-shadow:inset 0 12px 12px rgba(0, 0, 0, .5);
box-shadow:inset 0 12px 12px rgba(0, 0, 0, .5);
-webkit-border-top-right-radius: 295px;
-webkit-border-top-left-radius: 95px;
-moz-border-radius-topright: 95px;
-moz-border-radius-topleft: 95px;
border-top-right-radius: 95px;
border-top-left-radius: 95px;
}

    </style>
  </head>
  <body>
	<h1>{{.Title}}</h1>
    <p>
    	<ul>
    		<li><strong>IP</strong>: {{.Ip}}</li>
			<li><strong>MAC</strong>: {{.Mac}}</li>
    		<li><strong>RequestIp</strong>: {{.Request.RemoteAddr}}</li>
    		<li><strong>RequestURI</strong>: {{.Request.RequestURI}}</li>
    		{{range $key, $value := .Request.Header}} 
    			<li><strong>{{ $key }}</strong>: {{ $value }}</li>
    		{{end}}
 		<ul>
    </p>
    <div id="r-two-d-two-space">
	<div id="r-two-d-two">
	<div id="r-two-d-two-head-backunit"></div>
	<div id="r-two-d-two-head">
	<div id="r-two-d-two-head-top"><div id="r-two-d-two-head-topin"></div></div>
	<div id="r-two-d-two-head-projector-bluebox"></div>
	<div id="r-two-d-two-head-projector-bulb"></div>
	<div id="r-two-d-two-head-sensor">
	</div>
	<div class="r-two-d-two-head-sensor"></div>
	<div class="r-two-d-two-head-sensor"></div>
	<div id="r-two-d-two-head-bigsensor"></div>
	<div id="r-two-d-two-head-sensor-backhollow"></div>
	<div id="r-two-d-two-head-sensor-back"></div>
	<div id="r-two-d-two-head-sheen"></div>
	<div id="r-two-d-two-head-sensorband"></div>
	</div>
	<div id="r-two-d-two-body">
	<div id="r-two-d-two-body-back-one"></div>
	<div id="r-two-d-two-body-back-two"></div>
	<div id="r-two-d-two-body-back-three">
	<div class="r-two-d-two-body-back-threein"></div>
	<div class="r-two-d-two-body-back-threein"></div><div class="r-two-d-two-body-back-threein"></div></div>
	<div id="r-two-d-two-body-back-four"></div>
	<div class="r-two-d-two-body-sensor"></div>
	<div class="r-two-d-two-body-sensor"></div>
	<div class="r-two-d-two-body-sensor"></div>
	<div class="r-two-d-two-body-lowersensor"></div>
	<div class="r-two-d-two-body-lowersensor"></div>
	<div id="r-two-d-two-body-lowersensor-midbox">
	<div class="r-two-d-two-body-lowersensor-mid"></div>
	<div class="r-two-d-two-body-lowersensor-mid"></div>
	<div class="r-two-d-two-body-lowersensor-mid"></div>
	<div class="r-two-d-two-body-lowersensor-mid"></div>
	</div>
	<div id="r-two-d-two-body-bigspace"></div>
	<div id="r-two-d-two-body-sheen"></div>
	</div>
	<div id="r-two-d-two-trim"></div>
	<div id="r-two-d-two-leg-back"></div>
	<div id="r-two-d-two-foot-back"></div>
	<div id="r-two-d-two-lowerfoot-back"></div>
	<div id="r-two-d-two-leg-out"><div id="r-two-d-two-leg-top"></div><div id="r-two-d-two-leg-front"><div id="r-two-d-two-innerleg-top"><div id="r-two-d-two-leg-top-square"></div></div><div id="r-two-d-two-innerleg-front"><div id="r-two-d-two-innerleg-front-pipe"><div id="r-two-d-two-innerleg-front-pipein"></div></div></div><div id="r-two-d-two-leg-ankle"><div id="r-two-d-two-leg-anklebar"></div><div id="r-two-d-two-leg-anklein"></div></div></div>
	<div id="r-two-d-two-foot-front"></div>
	<div id="r-two-d-two-lowerfoot-round"></div>
	<div id="r-two-d-two-lowerfoot-front"></div>
	</div>
	</div>
	</div>
  </body>
</html>`

var tmplBack = `<html>
  <head>
    <title>{{.Title}}</title>
  </head>
  <body style="background-color:#{{printf "%x" .BgColor}};color:#{{printf "%x" .FgColor}}">
	<h1>{{.Title}}</h1>
    <p>
    	<ul>
    		<li><strong>IP</strong>: {{.Ip}}</li>
			<li><strong>MAC</strong>: {{.Mac}}</li>
    		<li><strong>RequestIp</strong>: {{.Request.RemoteAddr}}</li>
    		<li><strong>RequestURI</strong>: {{.Request.RequestURI}}</li>
    		{{range $key, $value := .Request.Header}}
    			<li><strong>{{ $key }}</strong>: {{ $value }}</li>
    		{{end}}
 		<ul>
    </p>
  </body>
</html>`

func handler(w http.ResponseWriter, r *http.Request) {
	requestIp, _, _ := net.SplitHostPort(r.RemoteAddr)

	fmt.Print(time.Now())
	fmt.Print(" - ")
	fmt.Print(requestIp)
	fmt.Print(" - ")
	fmt.Println(r.RequestURI)

	tmpl := tmplBack
	if strings.Contains(r.RequestURI, "graph") {
		tmpl = tmplGraph
	}
	t, err := template.New("foo").Parse(tmpl)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
	ip, macAddress, err := externalIP()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
	hashCode := int(crc32.ChecksumIEEE([]byte(ip + macAddress)))
	mask := 0x1000000
	bgcolor := hashCode%mask
	fgcolor := (mask - bgcolor -1)%mask

	p := &Page{Title: "Hello, World!", Ip: ip, Mac: macAddress, BgColor: bgcolor, FgColor: fgcolor, Request: r}
	err = t.Execute(w, p)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
}

func externalIP() (string, string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return "", "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), iface.HardwareAddr.String(), nil
		}
	}
	return "", "", errors.New("are you connected to the network?")
}

func main() {
	fmt.Println("Server Starting ...")
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(os.Getenv("HELLO_PORT"), nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("By!")
}