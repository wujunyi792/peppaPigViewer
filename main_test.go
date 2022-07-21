package main

import (
	"context"
	"fmt"
	browser "github.com/EDDYCJY/fake-useragent"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/time/rate"
	"log"
	"newJwCourseHelper/internal/config"
	"newJwCourseHelper/internal/core"
	"strings"
	"testing"
	"time"
)

func TestNewModule(t *testing.T) {
	res, err := core.LoadConfig(config.Config{
		Target:     []string{},
		ErrTag:     []string{},
		BucketFull: 5,
		Rate:       3,
		Ua:         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36",
	}).LoginPW("", "")
	if err != nil {
		panic(err)
	}
	//res.PrintCourseChosenList()
	res.SetTarget([]string{"(2022-2023-1)-A0104341-2"}).FindCourse().PrintFireCourseList()
	courses, err := res.FireCourses()
	if err != nil {
		panic(err)
	}
	fmt.Println(courses)
}

func TestRate(t *testing.T) {
	r := rate.NewLimiter(rate.Every(1*time.Second), 10)
	for {
		err := r.Wait(context.Background())
		if err != nil {
			panic(err)
			return
		}
		log.Println("getOne - ", r.Burst())
	}
}

func TestParse(T *testing.T) {
	body := `
<!doctype html>
<html lang="zh-CN">



<head>
	<title>&nbsp;</title>
	








<meta http-equiv="X-UA-Compatible" content="IE=edge" />
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
<meta name="Copyright" content="zfsoft" />	
<link rel="icon" href="/jwglxt/logo/favicon.ico?t=1656867463585" type="image/x-icon" />
<link rel="shortcut icon" href="/jwglxt/logo/favicon.ico?t=1656867463585" type="image/x-icon" />
<style type="text/css">	
	.active{font-weight: bolder;}
	#navbar-tabs li{ margin-top: 2px;}
	#navbar-tabs li a{ border-top: 2px solid transparent;}
	#navbar-tabs li.active a{border-top: 2px solid #0770cd;}
</style>


	
<!--jQuery核心框架库 -->
<script type="text/javascript" src="/zftal-ui-v5-1.0.2/assets/js/other_jquery/jquery.min.js?ver=27604308"></script>
<script type="text/javascript" src="/zftal-ui-v5-1.0.2/assets/js/jquery-migrate.min.js?ver=27604308"></script>
<!--jQuery浏览器检测 -->
<script type="text/javascript" src="/jwglxt/js/browse/browse-judge.js?ver=27604308"></script>
<script type="text/javascript">
	//浏览器版本验证
	var broswer = broswer();
	if(broswer.msie==true||broswer.safari==true||broswer.mozilla==true||broswer.chrome==true){
		if(broswer.msie==true&&broswer.version<9){
		   window.location.href = _path+"/xtgl/init_cxBrowser.html";
		}
	}else{
		 window.location.href = _path+"/xtgl/init_cxBrowser.html";
	}
</script>
<script type="text/javascript">
var _path 				= "/jwglxt";
var _systemPath 		= "/jwglxt";
var _stylePath 			= "/zftal-ui-v5-1.0.2";
var _reportPath 		= "http://jwbb.hdu.edu.cn/WebReport";
var _refreshInterval	= "10";
var _localeKey 			= "zh_CN";
</script>
<!--
jquery.ui 需要在bootstrap之前加载，解决冲突问题
http://blog.csdn.net/remote_roamer/article/details/14105999
-->
<script type="text/javascript" src="/zftal-ui-v5-1.0.2/assets/js/jquery-ui-custom.contact-1.0.0.js?ver=27604308"></script>
<!--Bootstrap布局框架-->
<link rel="stylesheet" type="text/css" href="/zftal-ui-v5-1.0.2/assets/plugins/bootstrap/css/bootstrap.min.css?ver=27604308" />
<link rel="stylesheet" type="text/css" href="/zftal-ui-v5-1.0.2/assets/css/error.css?ver=27604308" />
<link rel="stylesheet" type="text/css" href="/zftal-ui-v5-1.0.2/assets/css/zftal-ui.css?ver=27604308" />
<script type="text/javascript" src="/zftal-ui-v5-1.0.2/assets/plugins/bootstrap/js/bootstrap.min.js?ver=27604308" charset="utf-8"></script>
<!--jQuery常用工具扩展库：基础工具,资源加载工具,元素尺寸相关工具 -->
<script type="text/javascript" src="/zftal-ui-v5-1.0.2/assets/js/zftal/jquery.utils.contact-min.js?ver=27604308" charset="utf-8"></script>
<!--jQuery基础工具库：$.browser,$.cookie,$.actual等 -->
<script type="text/javascript" src="/zftal-ui-v5-1.0.2/assets/js/zftal/jquery.plugins.contact-min.js?ver=27604308" charset="utf-8"></script>
<!--jQuery自定义event事件库 -->
<script type="text/javascript" src="/zftal-ui-v5-1.0.2/assets/js/zftal/jquery.events.contact-min.js?ver=27604308" charset="utf-8"></script>
<!--JavaScript对象扩展库：Array,Date,Number,String -->
<script type="text/javascript" src="/zftal-ui-v5-1.0.2/assets/js/zftal/jquery.extends.contact-min.js?ver=27604308" charset="utf-8"></script>
<!--Bootbox弹窗插件-->
<link rel="stylesheet" type="text/css" href="/zftal-ui-v5-1.0.2/assets/plugins/bootbox/css/bootbox.css?ver=27604308" />
<script src="/zftal-ui-v5-1.0.2/assets/plugins/bootbox/bootbox.concat-min.js?ver=27604308" type="text/javascript" charset="utf-8"></script>
<script src="/zftal-ui-v5-1.0.2/assets/plugins/bootbox/lang/zh_CN.js?ver=27604308" type="text/javascript" charset="utf-8"></script>

<!--jQuery模拟滚动条库-->
<link rel="stylesheet" type="text/css" href="/zftal-ui-v5-1.0.2/assets/plugins/customscrollbar/css/jquery.mCustomScrollbar.min.css?ver=27604308" />
<script type="text/javascript" src="/zftal-ui-v5-1.0.2/assets/plugins/customscrollbar/js/jquery.mCustomScrollbar.min.js?ver=27604308" charset="utf-8"></script>
<!--jQuery.chosen美化插件-->
<link rel="stylesheet" type="text/css" href="/zftal-ui-v5-1.0.2/assets/plugins/chosen/css/chosen-min.css?ver=27604308" />
<script type="text/javascript" src="/zftal-ui-v5-1.0.2/assets/plugins/chosen/jquery.choosen.concat-min.js?ver=27604308" charset="utf-8"></script>
<script type="text/javascript" src="/zftal-ui-v5-1.0.2/assets/plugins/chosen/lang/zh_CN-min.js?ver=27604308" charset="utf-8"></script>
<script type="text/javascript" src="/zftal-ui-v5-1.0.2/assets/js/utils/jquery.utils.pinyin.min.js?ver=27604308" charset="utf-8"></script>
<!--[if lt IE 9]>
<script type="text/javascript" src="/zftal-ui-v5-1.0.2/assets/js/html5shiv.min.js?ver=27604308" charset="utf-8"></script>
<script type="text/javascript" src="/zftal-ui-v5-1.0.2/assets/js/respond.min.js?ver=27604308" charset="utf-8"></script>
<![endif]-->
<!--业务框架jQuery全局设置和通用函数库-->
<script type="text/javascript" src="/jwglxt/js/jquery.zftal.contact-min.js?ver=27604308"></script>
<!--国际化js库 -->
<script type="text/javascript" src="/zftal-ui-v5-1.0.2/assets/plugins/i18n/jquery.i18n-min.js?ver=27604308" charset="utf-8"></script>
<!--全局国际化js. -->
<script type="text/javascript" src="/jwglxt/js/globalweb/i18n-global_zh_CN.js?ver=27604308" charset="utf-8"></script>
<!--业务框架前端脚本国际化库-->
<script type="text/javascript" src="/zftal-ui-v5-1.0.2/assets/js/zftal/lang/jquery.zftal_zh_CN-min.js?ver=27604308"></script>
<!--密码强弱判断-->
<script type="text/javascript" src="/zftal-ui-v5-1.0.2/assets/js/utils/jquery.utils.strength.min.js?ver=27604308"></script>

	
	





	
<script type="text/javascript">

	var _path 		= "/jwglxt";
	var _systemPath = "/jwglxt";
	var _stylePath  = "/zftal-ui-v5-1.0.2";
	var _reportPath = "http://jwbb.hdu.edu.cn/WebReport";
	var _localeKey 			= "zh_CN";
	
	jQuery(function($){
		$('[data-toggle*="validation"]').trigger("validation");
		$('[data-toggle*="fixed"]').trigger("fixed");
		if($.fn.tooltip){
			$('[data-toggle*="tooltip"]').tooltip({container:'body'});
		}
	});
	
</script>
<style type="text/css">
	.captcha_modal{
		width: 380px;
		height: 330px;
		z-index: 9999;
		top: 150px;
		margin: auto;
		position: absolute;
		box-sizing: border-box;
		border-radius: 2px;
		background-color: #fff;
		box-shadow: 0 0 10px rgba(0,0,0,.3);
		left: 40%;
	}
</style>
<!-- 文件操作相关js -->
<script type="text/javascript" src="/jwglxt/js/globalweb/comp/file/file.js?ver=27604308"></script>
<!--教务系统通用业务js引用:比如学年，学期等公共的信息会放在这里-->
<script type="text/javascript" src="/jwglxt/js/globalweb/comp/i18n/jwglxt-common_zh_CN.js?ver=27604308"></script>
<!--业务模块的properties初始化-->
<!--国际化js库 -->
<script type="text/javascript" src="/jwglxt/js\globalweb\comp\i18n\N253512_zh_CN.js?ver=20210818" charset="utf-8"></script>

	<style type="text/css">
		.btn-quick{
			margin-right:30px;
			height:25px;
			padding: 0px 10px;
		}
		.outer_left .glyphicon:before {
		    margin: 14px;
		}
		.jxb-wyl{
			background-color:#f400006e !important;
		}
	</style>

</head>
<body>
<input type="hidden" id="pkey" name="pkey" value="" />
<input type="hidden" id="shyjsfbt" name="shyjsfbt" value="0" />
<input type="hidden" id="localeKey" name="localeKey" value="zh_CN" />
<input type="hidden" id="csrftoken" name="csrftoken" value="b0280804-c645-4b0f-a565-3b1190b411ad,b0280804c6454b0fa5653b1190b411ad"/>
<input type="hidden" id="cdTsxx" name="cdTsxx" value="-zfsplit-" />
<input type="hidden" id="wjylSfkf" name="wjylSfkf" value="0" >
	<!-- 头部 开始 -->
	<header class="navbar-inverse top2">
		<div class="container" id="navbar_container">
					<div class="container">
			<div class="navbar-header">
				<button class="navbar-toggle" type="button" data-toggle="collapse" data-target=".bs-navbar-collapse">
					<span class="sr-only"> 自主选课</span> 
					<span class="icon-bar"></span> 
					<span class="icon-bar"></span> 
					<span class="icon-bar"></span>
				</button>
				<a href="#" id="topButton" class="navbar-brand" onclick="onClickMenu('/xsxk/zzxkyzb_cxZzxkYzbIndex.html','N253512')">
					自主选课
				</a>
				<script type="text/javascript">
					document.title="自主选课";
				</script>
			</div>
		</div>
<!-- navbar-end  -->
		</div>
	</header>
	<script>
		if(window.self !== window.top){
			 $('body').css({
				"background": "#fff"
			}) 
			$('body').find('.navbar-inverse').hide();			
		}
	</script>
	
	<!--头部 结束 -->
	<div style="width: 100%; padding: 0px; margin: 0px;" id="bodyContainer">
		<!-- requestMap中的参数为系统级别控制参数，请勿删除 -->
		<form id="requestMap">
			 <input type="hidden" id="sessionUserKey" value="20081131" /> 
			 
			 	<input type="hidden" id="gnmkdmKey" value="N253512" />
			 
			 
		</form>
		<div class="container container-func sl_all_bg" id="yhgnPage">
			<div id="innerContainer">
				<!-- 放置页面显示内容 -->
				
	<div class="row sl_add_btn">
	    <div id="btn-groups" class="col-sm-12">
	    	<!-- 加载当前菜单栏目下操作   -->
			
			<!-- 加载当前菜单栏目下操作 -->
	    </div>
	</div>
	<div id="searchBox"></div> 
	
 	
 	
 	
 	
 		<div class="col-md-12 col-sm-12 border-b"  style="padding:8px 0px;">
			<div  style="float:left;padding:10px 15px;">
				<h5>
				<font id="xkxn"></font> 学年 <font id="xkxq"></font> 学期<font id="txt_xklc"></font><span id="sysj"></span>
				&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<b>本学期选课要求</b>总学分(不包括)最低&nbsp;<font color="red">16</font>
				
					&nbsp;&nbsp;最高&nbsp;<font color="red">36</font>
				
				
				&nbsp;&nbsp;&nbsp;本学期已选学分&nbsp;&nbsp;<font color="red" id="yxxfs">0</font>
				</h5>
			</div>
			<div style="margin-right:20px;float:right;padding:8px 15px;">
				<!-- <div style="float:left;margin-top:-2px;margin-right:20px;font-size:20px">【<a style="text-decoration:underline;" href="javascript:void(0)" onclick="$.showDialog(_path+'/xjyj/xsxyqk_ckXsXyxxHtmlView.html','我的修业情况',$.extend({},viewConfig,{width: ($('#yhgnPage').innerWidth()-200)+'px'}));">我的修业情况</a>】</div> -->
				<div id="quickXk" style="float:left;"></div>
				<div style="float:left;">
					<p style="margin-top:4px;margin-right:5px;float:left;border:1px solid #BCE8F1;background-color:#D9EDF7;height:15px;width:30px;"></p>未选
				</div>
				<div style="float:left;margin-left:20px">
					<p style="margin-top:4px;margin-right:5px;float:left;border:1px solid #BCE8F1;background-color:#fff7b2;height:15px;width:30px;"></p>重修未选
				</div>
				<div style="float:left;margin-left:20px">
					<p style="margin-top:4px;margin-right:5px;float:left;border:1px solid #BCE8F1;background-color:#C1FFC1;height:15px;width:30px;"></p>已选
				</div>
			</div>
		</div>

		<input type="hidden" name="iskxk" id="iskxk" value="1"/>
		<input type="hidden" name="jgh_id" id="jgh_id"/>
		<input type="hidden" name="jzxkf" id="jzxkf" value="0"/>
		<input type="hidden" name="xkzgmc" id="xkzgmc" value="20"/>
		<input type="hidden" name="xkzgxf" id="xkzgxf" value="36"/>
		<input type="hidden" name="zkcs" id="zkcs" value="11"/>
		<input type="hidden" name="zxfs" id="zxfs" value="21.8"/>
		<input type="hidden" name="bdzcbj" id="bdzcbj" value="2"/>
		<input type="hidden" name="xkxnm" id="xkxnm" value="2022"/>
		<input type="hidden" name="xkxqm" id="xkxqm" value="3"/>
		<input type="hidden" name="xkxnmc" id="xkxnmc" value="2022-2023"/>
		<input type="hidden" name="xkxqmc" id="xkxqmc" value="1"/>
		<input type="hidden" name="xh_id" id="xh_id" value="20081131"/>
		<input type="hidden" name="xqh_id" id="xqh_id" value="1"/>
		<input type="hidden" name="jg_id_1" id="jg_id_1" value="05"/>
		<input type="hidden" name="zyh_id" id="zyh_id" value="0523"/>
		<input type="hidden" name="zymc" id="zymc" value="计算机科学与技术"/>
		<input type="hidden" name="zyfx_id" id="zyfx_id" value="wfx"/>
		<input type="hidden" name="njdm_id" id="njdm_id" value="2020"/>
		<input type="hidden" name="njmc" id="njmc" value="2020"/>
		<input type="hidden" name="bh_id" id="bh_id" value="20052315"/>
		<input type="hidden" name="xbm" id="xbm" value="1"/>
		<input type="hidden" name="zh" id="zh" value=""/>
		<input type="hidden" name="xslbdm" id="xslbdm" value="7"/>
		<input type="hidden" name="mzm" id="mzm" value="01"/>
		<input type="hidden" name="xz" id="xz" value="4"/>
		<input type="hidden" name="ccdm" id="ccdm" value="8"/>
		<input type="hidden" name="xsbj" id="xsbj" value="4294967296"/>
		<input type="hidden" name="sjhm" id="sjhm" value="18149770580"/>
		<input type="hidden" name="xszxzt" id="xszxzt" value="1"/>
		<input type="hidden" name="njdm_id_1" id="njdm_id_1" value="2020"/>
		<input type="hidden" name="zyh_id_1" id="zyh_id_1" value="0523"/>
		<input type="hidden" name="sfxsxkbz" id="sfxsxkbz" value="1"/>
		<input type="hidden" name="sfxskssj" id="sfxskssj" value="0"/>
		<input type="hidden" name="wrljxbbhkg" id="wrljxbbhkg" value="0"/>
		<input type="hidden" name="jxbzbkg" id="jxbzbkg" value="1"/>
		<input type="hidden" name="tykpzykg" id="tykpzykg" value="0"/>
		<input type="hidden" name="tkdxyzms" id="tkdxyzms" value="0"/>
		<input type="hidden" name="jxbzhkg" id="jxbzhkg" value="0"/>
		<input type="hidden" name="xxdm" id="xxdm" value="10336"/>
		<input type="hidden" name="xkgwckg" id="xkgwckg" value="0"/>
		<input type="hidden" name="cxkctskg" id="cxkctskg" value="0"/>
		<input type="hidden" name="kxqxktskg" id="kxqxktskg" value="0"/>
		<input type="hidden" name="tbtkxqxktskg" id=tbtkxqxktskg value="0"/>
		<input type="hidden" name="xkkczdsqkg" id="xkkczdsqkg" value="1"/>
		<input type="hidden" name="xkmcjzxskcs" id="xkmcjzxskcs" value="10"/>
		
		
		<input type="hidden" name="zzxkgjcxkg_kcz" id="zzxkgjcxkg_kcz" value="0"/>
		<input type="hidden" name="zzxkgjcxkg_nj" id="zzxkgjcxkg_nj" value="1"/>
		<input type="hidden" name="zzxkgjcxkg_xy" id="zzxkgjcxkg_xy" value="1"/>
		<input type="hidden" name="zzxkgjcxkg_zy" id="zzxkgjcxkg_zy" value="1"/>
		<input type="hidden" name="zzxkgjcxkg_kkxy" id="zzxkgjcxkg_kkxy" value="1"/>
		<input type="hidden" name="zzxkgjcxkg_xqu" id="zzxkgjcxkg_xqu" value="0"/>
		<input type="hidden" name="zzxkgjcxkg_yqu" id="zzxkgjcxkg_yqu" value="0"/>
		<input type="hidden" name="zzxkgjcxkg_tjbj" id="zzxkgjcxkg_tjbj" value="0"/>
		<input type="hidden" name="zzxkgjcxkg_kclb" id="zzxkgjcxkg_kclb" value="1"/>
		<input type="hidden" name="zzxkgjcxkg_kcxz" id="zzxkgjcxkg_kcxz" value="1"/>
		<input type="hidden" name="zzxkgjcxkg_jxms" id="zzxkgjcxkg_jxms" value="1"/>
		<input type="hidden" name="zzxkgjcxkg_kcgs" id="zzxkgjcxkg_kcgs" value="1"/>
		<input type="hidden" name="zzxkgjcxkg_skxq" id="zzxkgjcxkg_skxq" value="1"/>
		<input type="hidden" name="zzxkgjcxkg_skjc" id="zzxkgjcxkg_skjc" value="1"/>
		<input type="hidden" name="zzxkgjcxkg_jxb" id="zzxkgjcxkg_jxb" value="1"/>
		<input type="hidden" name="zzxkgjcxkg_sfcx" id="zzxkgjcxkg_sfcx" value="1"/>
		<input type="hidden" name="zzxkgjcxkg_ywyl" id="zzxkgjcxkg_ywyl" value="1"/>
		<input type="hidden" name="zzxkgjcxkg_sksjct" id="zzxkgjcxkg_sksjct" value="0"/>
		
		
		<input type="hidden" name="xkczbj" id="xkczbj" value="1"/>
		<input type="hidden" name="isSlct" id="isSlct" value="0"/>
		<input type="hidden" name="kklxdm" id="kklxdm" value=""/>
		<input type="hidden" name="xkkz_id" id="xkkz_id" value=""/>
		<input type="hidden" name="jxbzb" id="jxbzb" value=""/>
		
		
			<div class="panel panel-info">
			<!-- 开始 -->
			<ul class="nav nav-tabs sl_nav_tabs" role="tablist" id="nav_tab">
				
					
						<li class="active"><a href="javascript:void(0)" onclick="queryCourse(this,'01','E0BE1EB065FBFA29E0536264A8C04A31','2020','0523')" role="tab" data-toggle="tab">主修课程</a></li>
						<input type="hidden" name="firstKklxdm" id="firstKklxdm" value="01"/>
						<input type="hidden" name="firstXkkzId" id="firstXkkzId" value="E0BE1EB065FBFA29E0536264A8C04A31"/>
						<input type="hidden" name="firstNjdmId" id="firstNjdmId" value="2020"/>
						<input type="hidden" name="firstZyhId" id="firstZyhId" value="0523"/>
			 		
			 		
		 		
					
			 		
			 			<li><a href="javascript:void(0)" onclick="queryCourse(this,'10','E0BDC4C7604BD44BE0536264A8C0B7EC','2020','0523')" role="tab" data-toggle="tab">通识选修课</a></li>
			 		
		 		
					
			 		
			 			<li><a href="javascript:void(0)" onclick="queryCourse(this,'05','E0BE43551AEB415FE0536164A8C06426','2020','0523')" role="tab" data-toggle="tab">体育分项</a></li>
			 		
		 		
					
			 		
			 			<li><a href="javascript:void(0)" onclick="queryCourse(this,'09','E0BE43551B64415FE0536164A8C06426','2020','0523')" role="tab" data-toggle="tab">特殊课程</a></li>
			 		
		 		
					
			 		
			 			<li><a href="javascript:void(0)" onclick="queryCourse(this,'08','E17228E7DE048AE4E0536264A8C0F850','2020','0523')" role="tab" data-toggle="tab">重修课程</a></li>
			 		
		 		
		 		<div class="pull-right" style="margin-top:4px;margin-right:30px">
		 			 
		 		</div>
			</ul>
			</div>
		
	
	<div id="displayBox"></div>
	<div id="choosedBox"></div>
	<div id="endsign" style="display:none; text-align:center; height: 50px"><i class="red">......已到最后......</i></div><!-- （共 <font id="searchCount"></font> 条记录） -->
	<!-- <div id="waitsign" style="display:none; text-align:center; height: 50px"><i class="red bigger-300 icon-spinner icon-spin"></i></div> -->
	<div id="more" style="text-align:center; display:none"><font color="#2a6496" size="5">[<a href="javascript:void(0)" onclick="loadCoursesByPaged();">点此查看更多</a>]</font></div>
	<!--jQuery.jqGrid -->
<link rel="stylesheet" type="text/css" href="/zftal-ui-v5-1.0.2/assets/plugins/jqGrid/css/jquery.jqgrid.css?ver=27604308" />
<script type="text/javascript" src="/zftal-ui-v5-1.0.2/assets/plugins/jqGrid/jquery.jqgrid.src-min.js?ver=27604308" charset="utf-8"></script>
<script type="text/javascript" src="/zftal-ui-v5-1.0.2/assets/plugins/jqGrid/jquery.jqgrid.contact-min.js?ver=27604308" charset="utf-8"></script>
<script type="text/javascript" src="/zftal-ui-v5-1.0.2/assets/plugins/jqGrid/lang/zh_CN.js?ver=27604308" charset="utf-8"></script>
<script type="text/javascript" src="/jwglxt/js/plugins/jqGrid4.6/jquery.jqgrid.settings.js?ver=27604308" charset="utf-8"></script>
<script type="text/javascript" src="/jwglxt/js/plugins/jqGrid4.6/jquery.jqGrid-min.js?ver=27604308" charset="utf-8"></script>
	<!--jQuery.validation -->
<link rel="stylesheet" type="text/css" href="/zftal-ui-v5-1.0.2/assets/plugins/validation/css/jquery.validate-min.css?ver=27604308" />
<script type="text/javascript" src="/zftal-ui-v5-1.0.2/assets/plugins/validation/js/jquery.validate-min.js?ver=27604308" charset="utf-8"></script>
<script type="text/javascript" src="/zftal-ui-v5-1.0.2/assets/plugins/validation/js/jquery.validate.contact-min.js?ver=27604308" charset="utf-8"></script>
<script type="text/javascript" src="/zftal-ui-v5-1.0.2/assets/plugins/validation/js/jquery.validate.methods.contact-min.js?ver=27604308" charset="utf-8"></script>
<script type="text/javascript" src="/zftal-ui-v5-1.0.2/assets/plugins/validation/lang/zh_CN-min.js?ver=27604308" charset="utf-8"></script>
<script type="text/javascript">
	jQuery(function($){
		$('[data-toggle*="validation"]').trigger("validation");
	});
</script>

	<link rel="stylesheet" type="text/css" href="/jwglxt/js/plugins/searchbox/css/jquery.searchbox-min.css?ver=27604308" />
<script type="text/javascript" src="/jwglxt/js/plugins/searchbox/jquery.searchbox.contact-min.js?ver=27604308" charset="utf-8"></script>
<script type="text/javascript" src="/jwglxt/js/plugins/searchbox/jquery.searchbox.contact-zh_CN.js?ver=27604308" charset="utf-8"></script>
	<script type="text/javascript" src="/jwglxt/js/comp/jwglxt/xkgl/xsxk/zzxkYzb.js?ver=27604308"></script>

				




			</div>
		</div>
	</div>
	<!-- footer -->
	







<!-- footer --> 

<div id="footerID" class="footer"  style="background-color: " >
	
	<p>版权所有&#169; Copyright 1999-2022 正方软件股份有限公司　　中国·杭州西湖区紫霞街176号 互联网创新创业园2号301&nbsp;&nbsp;&nbsp;版本V-8.0.0</p>
</div>




<script  type="text/javascript">
	//系统中页面底部的时间随系统时间来定
	/* var date=new Date;
	var year=date.getFullYear();
	var s = $("#footerID").text().indexOf('-');
	var textNr =$("#footerID").text().substr(0,s+1)+year+$("#footerID").text().substr(s+5);
	$("#footerID").text("");
	$("#footerID").text(textNr); */

</script>
<!-- footer-end -->
	<!-- footer-end -->
</body>
<script type="text/javascript">
	jQuery(function($) {	
		if ($("#navbar-tabs").length > 0) {
			$("#navbar-tabs li:eq(0) a").tab('show');
		}
		setMinheight();
		setCdtsxx();
		
		$('#yhgnPage').resize(function(){
			var minHeight  = $(this).data('min-height');
			var innerHeight =  $('#innerContainer').outerHeight(true);
			if(self != top){
				$(this).css('min-height', "380");
			}else{
				$(this).css('min-height', Math.max(minHeight,innerHeight));
			}
			
		});
		
		function setMinheight(){
			//计算grid页面高度
			var docuemntHeight = $(document).height();
			var topHeight = $('header').height();
			var footerHeight = $('.footer').height();
			var containerHeight = docuemntHeight-topHeight-footerHeight-60;
			if(self != top){
				containerHeight = 380;
			}
			$('#yhgnPage').css('min-height',containerHeight).data('min-height',containerHeight);
		}
		
		//加载title名称，根据系统内置表设置，取菜单名称或者系统名称
		$(document).attr("title",'自主选课');
		
		//设置菜单提示信息
		function setCdtsxx(){
			if($("#cdTsxx").val()!=null&&$("#cdTsxx").val()!="-zfsplit-"){
				var tsxx = $("#cdTsxx").val().split("-zfsplit-")[0];
				var tsxxljdz = $("#cdTsxx").val().split("-zfsplit-")[1];
				var tsxxHtml = "<a class='navbar-brand' "+(tsxxljdz!=""?"href='"+tsxxljdz+"' target='_blank'":"")+">"+tsxx+"</a>";
				$(".navbar-header").append(tsxxHtml);
			}
		}
	});
</script>
<!-- 设置元素固定位置显示插件  scrolltofixed 相应ini文件引用-->
<!-- è®¾ç½®åç´ åºå®ä½ç½®æ¾ç¤ºæä»¶  scrolltofixed -->
<script type="text/javascript" src="/jwglxt/js/plugins/scrolltofixed/jquery-scrolltofixed-min.js?ver=27604308"></script>
<script>
	$(document).ready(function() {
		$("div.sl_add_btn .btn-toolbar").scrollToFixed({
			marginTop:35,
			zIndex:1050,
			fixed:function(){
				$(this).css("width","auto");
			}
		});
		$("#gbox_tabGrid div.ui-state-default").scrollToFixed({
				//marginTop:3,
			    spacerClass:'hide-div',
		});
	});
</script>
<style>
	.hide-div{
		display:none; 
	}
</style>
<!-- 软件评价 相应ini文件引用-->


<link rel="stylesheet"  type="text/css" href="/jwglxt/js/plugins/tagtree/tagTree.css?ver=27604308"/>
<script type='text/javascript' src="/jwglxt/js/plugins/tagtree/tagtree.js?ver=27604308"></script>
<script type='text/javascript' src="/jwglxt/js/plugins/tagtree/tagtreeBusiness.js?ver=27604308"></script>

</html>`

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("input").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("name")
		value, _ := s.Attr("value")
		fmt.Printf("%d: %s %s\n", i, name, value)
	})
}

//func TestMakeForm(t *testing.T) {
//	form := FindClassReq{
//		FilterList: []string{"1", "2"},
//	}
//	println(form.makeForm())
//}

func TestUA(t *testing.T) {
	for i := 0; i < 100; i++ {
		println(browser.Random())
	}
}
