package parser

import (
	"testing"
)

func TestFindPagelets(t *testing.T) {
	test_html := `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xmlns:wb="http://open.weibo.com/wb">
	<head>
		<meta http-equiv="Content-type" content="text/html; charset=utf-8" />
		<title>邻居帮帮忙-瓜子社区_找到附近的人，解决身边的事</title>
		<meta name="keywords" content="回龙观,回龙观社区,回龙观社区网,回龙观网，回龙观论坛,天通苑,天通苑社区,天通苑社区网,天通苑网，天通苑论坛,望京,望京社区,望京社区网,望京网，望京论坛" />
		<meta name="description" content="瓜子是基于位置的社区网站，我们的宗旨是：找到身边的人，解决身边的事，欢迎使用瓜子社区！" />
		<meta content="IE=8" http-equiv="X-UA-Compatible"/>
		<script type="text/javascript">window.__GJ_PACK_CONFIG__ = {"documentDomain":"ganji.com","debug":false,"useCombine":true,"addVersion":true,"cookieDomain":"ganji.com","defaultServer":"sta.ganji.com","servers":["sta.ganjistatic1.com"],"fileVersions":{"3":1318347027,"5":1319637755,"7":1319194905,"9":1324465245,"11":1337314973,"15":1338793227,"19":1324293456,"20":1340285048,"21":1340781465,"22":1340703972,"23":1339470360,"24":1340285049,"25":1339142998,"26":1328261259,"27":1340781491,"30":1334834067,"32":1334834067,"34":1334834067,"36":1334834067,"38":1334834067,"40":1334834067,"41":1337314973,"42":1303201742,"44":1329736022,"46":1334909111,"48":1303201735,"50":1303201742,"52":1314092918,"54":1337314831,"56":1324461546,"58":1303201741,"60":1303201741,"62":1337942336,"64":1305019612,"66":1303201741,"68":1303201741,"70":1303201741,"71":1322213567,"73":1338285664,"74":1340620992,"75":1337864386,"77":1324523310,"81":1339040086,"82":1339038345,"83":1339038322,"85":1340007209,"87":1331279027,"88":1329139546,"90":1337314036,"92":1337937842,"94":1303201741,"95":1336710310,"97":1318347028,"98":1314102295,"100":1332416447,"101":1319637755,"103":1336710310,"104":1316516316,"106":1321429052,"107":1340954008,"110":1334305721,"113":1322805178,"115":1336710310,"116":1331026186,"117":1319720921,"120":1340865275,"121":1341224213,"123":1341224212,"125":1340865271,"126":1341224224,"127":1341224602,"129":1341224213,"130":1340865273,"132":1332751329,"134":1340173266,"136":1322794220,"137":1340173283,"138":1340948735,"139":1340948736,"140":1337758024,"142":1338949631,"144":1340871743,"145":1340865267,"146":1340948735,"147":1340865266,"148":1338949625,"149":1340173263,"151":1328671707,"153":1340865268,"154":1340173265,"155":1340865267,"156":1341224211,"157":1340865270,"160":1326092263,"161":1336714118,"163":1338949627,"165":1341196075,"166":1340173269,"167":1340865268,"168":1337758022,"170":1338966254,"172":1336378331,"173":1340865271,"176":1335323472,"179":1326092263,"180":1326092263,"181":1340173271},"fileCombines":{"11":[182,183,184],"27":[182,9,15,19,20,21,22,23,24,25,26],"41":[30,32,34,36,38,40],"127":[182,185,186,187,188,189],"126":[190,191,192],"137":[193,194]},"fileCodes":{"tool\/webim\/js\/webim.js":27,"ganji":83,"app_global":85,"autocomplete":87,"jquery":3,"jquery_ui":88,"util":11,"json":90,"log_tracker":92,"select":94,"upload":196,"iframe":[182,5],"panel":[95,183,7,9],"flash":[182,97,98],"talk_to_parent":[182,100,101],"validator":[182,103,104],"validator2":[182,103,106,107],"map":[183,110],"uploader":[183,113],"editor":[115,116],"passport":[117],"guazi":138,"gzLogin":139,"jqueryEasing":140,"gzValidator":142,"gzBaseCss":144,"gzPersonalCenterCss":145,"gzPanelCss":146,"gzCoverCss":147,"gzInfosCss":148,"gzProcessCss":149,"gzAction":121,"gzApi":120,"gzNameCardTip":125,"gzObservable":123,"gzAutoComplete":151,"gzFocusPic":153,"gzGallery":154,"gzCropImage":155,"gzPublisher":134,"gzSelectPanel":156,"gzFileUpload":157,"ajaxUpload":136,"gzMapAutoComplete":160,"gzMap":161,"gzTextareaUtil":130,"gzUtil":129,"gzPlaceHolder":132,"gzPrivateMsg":163,"gzPanel":165,"gzPassword":166,"gzEmoticon":167,"notice":168,"gzCommentBox":170,"gzBeautySelect":172,"gzAskForHelp":173,"gzGuide":176,"gzDatePicker":[182,179,180],"emailTip":181,"gzCmbBase":127,"gzCmbUtil":126,"gzCmbPublish":137},"words":["js\/","0|util\/","1|jquery\/","2|jquery-1.4.4.js","1|iframe\/","4|iframe.js","1|dragdrop\/","6|dragdrop.js","1|panel\/","8|panel.js","1|ganji\/","10|util.cmb.js","0|app\/","12|common\/","13|user\/","14|user.js","tool\/","16|webim\/","17|js\/","18|md5-min.js","18|webim.util.js","18|webim.core.js","18|webim.widget.js","18|webim.widget.mayi.js","18|webim.html.js","18|webim.start.js","18|webim.interfaces.js","18|webim.cmb.js","1|yui\/","28|yahoo-dom-event\/","29|yahoo-dom-event.js","28|json\/","31|json.js","28|element\/","33|element.js","28|datasource\/","35|datasource.js","28|swf\/","37|swf.js","28|charts\/","39|charts.js","28|charts.cmb.js","4|closeproxy.html","1|string\/","43|string.js","1|uuid\/","45|uuid.js","swf\/","47|cache.swf","1|datetime\/","49|active_time.js","13|lottery\/","51|lottery.js","51|image\/","53|song_img.gif","13|detail\/","55|detail.js","13|google\/","57|google.js","13|langtaojin\/","59|langtaojin.js","13|adm\/","61|adm.js","13|ppc\/","63|ppc.js","13|login\/","65|login.js","13|message\/","67|message.js","13|jubao\/","69|jubao.js","18|webim.js","17|css\/","72|remind.css","72|webim.css","72|webim_page.css","12|data\/","76|city_domain.js","app\/","78|common\/","79|uniontgm\/","80|uniontgm.js","80|uniontgm_tpl.js","10|ganji.js","13|global\/","84|global.js","1|autocomplete\/","86|autocomplete.js","2|jquery.ui.js","1|json\/","89|json.js","1|log_tracker\/","91|log_tracker_simple.js","1|select\/","93|jquery.select.js","8|panel.css","1|swfobject\/","96|swfobject-2.2.js","96|swfloader.js","1|window_name\/","99|window_name.js","4|talk_to_parent.js","1|validator\/","102|validator.css","102|validator.js","1|event\/","105|event.js","102|validator-2.js","16|map\/","108|js\/","109|create_map.js","16|uploader\/","111|js\/","112|create_uploader.js","16|editor\/","114|editor.css","114|editor.js","14|passport.js","78|guazi\/","118|js\/","119|guaziApi.js","119|guaziAction.js","119|observable\/","122|observable.js","119|nameCardTip\/","124|nameCardTip.js","118|util.cmb.js","118|base.cmb.js","119|util\/","128|util.js","128|textareaUtil.js","119|placeHolder\/","131|placeHolder.js","119|publisher\/","133|publisher.js","119|ajaxUpload\/","135|ajaxUpload.js","118|publish.cmb.js","118|guazi.base.js","128|login.js","128|jquery.easing.1.3.js","119|validator\/","141|validator.js","118|css\/","143|base.css","143|personalCenter.css","143|panel.css","143|cover.css","143|infos.css","143|process.css","119|autocomplete\/","150|autocomplete.js","119|picUtil\/","152|focusPic.js","152|gallery.js","152|cropImage.js","133|selectPanel.js","133|fileUpload.js","118|tool\/","158|map\/","159|mapAutoComplete.js","159|map.js","119|msg\/","162|privateMsg.js","119|panel\/","164|gz_panelB.js","128|password.js","133|emoticon.js","162|notice.js","119|comment\/","169|comment-box.js","119|beautySelect\/","171|beautySelect.js","162|askForHelp.js","119|actives\/","174|guide\/","175|gzGuide.js","119|datetime\/","177|date_picker\/","178|date_picker.css","178|date_picker.js","128|emailTip.js","jquery","iframe","panel","gzApi","gzAction","gzObservable","gzNameCardTip","gzCmbUtil","gzUtil","gzTextareaUtil","gzPlaceHolder","gzPublisher","ajaxUpload","1|upload\/","195|upload.js"],"lottery":false};</script>
<script src="http://sta.ganjistatic1.com/public/js/util/ganji/ganji.__1339038322__.js" type="text/javascript"></script>
<script src="http://sta.ganjistatic1.com/public/app/guazi/guazi.base.__1340948735__.js" type="text/javascript"></script>
<link href="http://sta.ganjistatic1.com/public/app/guazi/app/bbm/css/help.css" rel="stylesheet" type="text/css" />
	</head>
	<body>
		<div class="helpWrap">
			<div class="header">
				<a target="_blank" href="http://guazi.com" ><img src="http://sta.ganjistatic1.com/src/app/guazi/app/bbm/image/bg1.jpg" width="950" height="90" alt="瓜子社区" title="瓜子社区"/></a>
			</div>
			<div class="main clearfix">
				<ul class="main-nav">
					<li class="selected"><a href="/app/bbm/" class="home">首页</a></li>
					<li><a href="/app/bbmSetLoc?state=1" class="publishHelp" >发布求助</a></li>
					<li><a href="/app/bbmSetLoc?state=2" class="otherHelp" >我要帮忙</a></li>
										<li><a target="_blank" href="http://guazi.ganji.com/profile-27084473-332.html" class="myHelp">我的求助</a></li>
									</ul>
                <div class="main-cont">
					<div id="testDiv1" priority="1" href="http://ok.com" pagelet="true" />
					<p class="helpDetail-des">共 <span class="tred">587</span> 位邻居得到帮助</p>
					<div id="focus-pic-id" class="" style="position:relative;width:740px;height:280px;">
						<div class="helpDetail-tag tag1-on" id="helpDetail-tag">
							<span class="tag1 tags">狗狗找新主人</span>
							<span class="tag2 tags">找物业电话</span>
							<span class="tag3 tags">寻靠谱月嫂</span>
							<span class="tag4 tags">转让闲置鞋子</span>
							<span class="tag5 tags">找狗狗</span>
						</div>
						<div class="helpDetail" id="helpDetail">
							<div class="helpDetail-change" style="position:absolute;width:738px;left:0;top:0;">
								<div class="helpDetail-tit clearfix">
									<h2 style="margin-right:15px;"><a target="_blank" href="http://guazi.ganji.com/pb-53237.html" class="agrey">在瓜子上，我帮我的狗狗找到了新主人</a></h2>	
									<div class="box-tw2 clearfix" style="float:left;">
										<a target="_blank" href="http://guazi.ganji.com/profile/?uid=95958818" class="box-tw2-img"><img src="http://sta.ganjistatic1.com/src/app/guazi/app/bbm/image/head1.jpg" width="30" height="30" alt="小公主" title="小公主"/></a>
										<div class="box-tw2-cont tgrey3">
											<p><a target="_blank" href="http://guazi.ganji.com/profile/?uid=95958818" target="_blank">小公主</a></p>
											<p>居住地：<span class="tgrey2">龙禧苑</span></p>
										</div>
									</div>
								</div>
								<div class="helpDetail-cont">
									<span class="text1">2012.6.1<br>狗狗寻觅新主人</span>
									<span class="text2">2012.6.6<br>在瓜子发帖求助</span>
									<span class="text3">2012.6.8<br>找到主人</span>
									<span class="text4">历时<var>2</var>天</span>
									<span class="text5">穿越<var>7</var>个小区</span>
									<span class="text6">共有<var>33</var>位邻居帮忙</span>
									<a target="_blank" href="http://guazi.ganji.com/pb-53237.html" class="check" hidefocus="true">查看事件经过</a>
								</div>					
							</div>
							<div class="helpDetail-change" style="position:absolute;width:738px;left:0;top:0;display:none;">
								<div class="helpDetail-tit clearfix">
									<h2 style="margin-right:15px;"><a target="_blank" href="http://guazi.ganji.com/pb-53330.html" class="agrey">在瓜子上，邻居就帮我找到了小区的物业电话，真是太感谢了</a></h2>	
									<div class="box-tw2 clearfix" style="float:left;">
										<a target="_blank" href="http://guazi.ganji.com/profile/?uid=84184865" class="box-tw2-img"><img src="http://sta.ganjistatic1.com/src/app/guazi/app/bbm/image/head2.jpg" width="30" height="30" alt="嫩香鸡腿堡" title="嫩香鸡腿堡"/></a>
										<div class="box-tw2-cont tgrey3">
											<p><a target="_blank" href="http://guazi.ganji.com/profile/?uid=84184865" target="_blank">嫩香鸡腿堡</a></p>
											<p>居住地：<span class="tgrey2">和谐家园</span></p>
										</div>
									</div>
								</div>
								<div class="helpDetail-cont">
									<span class="text1">2012.6.7<br>寻觅小区物业</span>
									<span class="text2">2012.6.7<br>在瓜子发帖求助</span>
									<span class="text3">2012 6.7<br>找到了物业的电话</span>
									<span class="text4">历时<var>30</var>分钟</span>
									<span class="text5">穿越<var>5</var>个小区</span>
									<span class="text6">共有<var>15</var>位邻居帮忙</span>
									<a target="_blank" href="http://guazi.ganji.com/pb-53330.html" class="check" hidefocus="true">查看事件经过</a>
								</div>					
							</div>
							<div class="helpDetail-change" style="position:absolute;width:738px;left:0;top:0;display:none;">
								<div class="helpDetail-tit clearfix">
									<h2 style="margin-right:15px;"><a target="_blank" href="http://guazi.ganji.com/pb-47672.html" class="agrey">我是一个新妈妈，在瓜子上我给宝宝找到了合适的月嫂</a></h2>	
									<div class="box-tw2 clearfix" style="float:left;">
										<a target="_blank" href="http://guazi.ganji.com/profile/?uid=112417360" class="box-tw2-img"><img src="http://sta.ganjistatic1.com/src/app/guazi/app/bbm/image/head3.jpg" width="30" height="30" alt="牧扬猫" title="牧扬猫"/></a>
										<div class="box-tw2-cont tgrey3">
											<p><a target="_blank" href="http://guazi.ganji.com/profile/?uid=112417360" target="_blank">牧扬猫</a></p>
											<p>居住地：<span class="tgrey2">龙跃苑</span></p>
										</div>
									</div>
								</div>
								<div class="helpDetail-cont">
									<span class="text1">2012.5.11<br>寻找月嫂</span>
									<span class="text2">2012.5.14<br>在瓜子发帖求助</span>
									<span class="text3">2012.5.18<br>找到了月嫂</span>
									<span class="text4">历时<var>4</var>天</span>
									<span class="text5">穿越<var>5</var>个小区</span>
									<span class="text6">共有<var>8</var>位邻居帮忙</span>
									<a target="_blank" href="http://guazi.ganji.com/pb-47672.html" class="check" hidefocus="true">查看事件经过</a>
								</div>					
							</div>
							<div class="helpDetail-change" style="position:absolute;width:738px;left:0;top:0;display:none;">
								<div class="helpDetail-tit clearfix">
									<h2 style="margin-right:15px;"><a target="_blank" href="http://guazi.ganji.com/pb-39738.html" class="agrey">我是北北，我为我的买小了的小鞋子找到了新主人</a></h2>	
									<div class="box-tw2 clearfix" style="float:left;">
										<a target="_blank" href="http://guazi.ganji.com/profile/?uid=49881212" class="box-tw2-img"><img src="http://sta.ganjistatic1.com/src/app/guazi/app/bbm/image/head4.jpg" width="30" height="30" alt="北北" title="北北"/></a>
										<div class="box-tw2-cont tgrey3">
											<p><a target="_blank" href="http://guazi.ganji.com/profile/?uid=49881212" target="_blank">北北</a></p>
											<p>居住地：<span class="tgrey2">首开智慧社</span></p>
										</div>
									</div>
								</div>
								<div class="helpDetail-cont">
									<span class="text1">2012.4.9<br>鞋子买小了</span>
									<span class="text2">2012.4.11<br>在瓜子发帖求助</span>
									<span class="text3">2012.4.20<br>鞋子找到了新主人</span>
									<span class="text4">历时<var>9</var>天</span>
									<span class="text5">穿越<var>6</var>个小区</span>
									<span class="text6">共有<var>34</var>位邻居帮忙</span>
									<a target="_blank" href="http://guazi.ganji.com/pb-39738.html" class="check" hidefocus="true">查看事件经过</a>
								</div>				
							</div>
							<div class="helpDetail-change" style="position:absolute;width:738px;left:0;top:0;display:none;">
								<div class="helpDetail-tit clearfix">
									<h2 style="margin-right:15px;"><a target="_blank" href="http://guazi.ganji.com/pb/?pid=48382" class="agrey">我是双双，我帮瓜子邻居找到了走丢的小哈宝宝</a></h2>	
									<div class="box-tw2 clearfix" style="float:left;">
										<a target="_blank" href="http://guazi.ganji.com/profile/?uid=64365657" class="box-tw2-img"><img src="http://sta.ganjistatic1.com/src/app/guazi/app/bbm/image/head5.jpg" width="30" height="30" alt="双双" title="双双"/></a>
										<div class="box-tw2-cont tgrey3">
											<p><a target="_blank" href="http://guazi.ganji.com/profile/?uid=64365657" target="_blank">双双</a></p>
											<p>居住地：<span class="tgrey2">立水桥</span></p>
										</div>
									</div>
								</div>
								<div class="helpDetail-cont">
									<span class="text1">2012.5.15<br>邻居狗狗走失</span>
									<span class="text2">2012.5.16<br>在瓜子发帖求助</span>
									<span class="text3">2012.5.17<br>找到了狗狗</span>
									<span class="text4">历时<var>1</var>天</span>
									<span class="text5">穿越<var>10</var>个小区</span>
									<span class="text6">共有<var>60</var>位邻居帮忙</span>
									<a target="_blank" href="http://guazi.ganji.com/pb/?pid=48382" class="check" hidefocus="true">查看事件经过</a>
								</div>				
							</div>
						</div>
					</div>
					<p class="tac buttW1" style="margin-top:40px;">
						<a id="findButton" class="butt1 mgr10" href="javascript:;">我要发布求助，找邻居帮忙</a>
						<a id="helpButton" class="butt4" href="javascript:;">我要帮邻居</a>
					</p>
                </div>
			</div>
			<div id="testDiv2" priority="1" href="http://ok.com" pagelet="true" />
			<div class="footer clearfix">
				<p class="fl tgrey3">瓜子社区(赶集旗下社区网站)</p>
				<script src="http://tjs.sjs.sinajs.cn/open/api/js/wb.js?appkey=1222281205" type="text/javascript" charset="utf-8"></script>
				<p class="fr">
					<wb:follow-button uid="2635078750" type="red_1" width="63" height="24" style="float: left;"></wb:follow-button>
					<wb:publish button_size="small" button_text="分享给好友" height="24" default_text="#瓜子邻居帮帮忙# 终于有一个东东可以满足我的需要啦，我推荐【邻居帮帮忙】，好神奇啊！用ta能找到附近的邻居，向邻居求助，马上就有人来帮忙啦！方便使用，来试试吧~ http://guazi.ganji.com/app/bbm" style="float: left;">微博发布器</wb:publish>
				</p>
			</div>	
		</div>
		<script type="text/javascript">
            GJ.use('gzUtil , gzFocusPic' , function(){
                $('#findButton').click(function(){
                    GZ.sinaAuth('bbm', function(){
						window.location.href = "http://guazi.ganji.com/app/bbmSetLoc?state=1";
                    });
                });
				$('#helpButton').click(function(){
                    GZ.sinaAuth('bbm', function(){
                        window.location.href = "http://guazi.ganji.com/app/bbmSetLoc?state=2";
                    });
                });
                GJ.use('gzFocusPic' , function(){
                    GZ.focusPic({
                        FocusPicID: 'focus-pic-id',
                        BigPicID: 'helpDetail',
                        NumbersID: 'helpDetail-tag',
						time: 3000,
                        callback: function(index){
                            $('#helpDetail-tag').attr('class' , 'helpDetail-tag tag' + (index + 1) + '-on');
                        }
                    });
                });
            });
		</script>
		<div id="testDiv3" priority="1" href="http://ok.com" pagelet="true" />
	</body>
</html>
	`
	pagelets := FindPagelets(test_html)
	for i := 0; i < 3; i++ {
		println(pagelets[i].id)
		println(pagelets[i].priority)
		println(pagelets[i].href)
	}
}
