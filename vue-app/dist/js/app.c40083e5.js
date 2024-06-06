(function(){"use strict";var e={1627:function(e,o,t){var r=t(5130),a=(t(8077),t(6768));function l(e,o,t,r,l,s){const n=(0,a.g2)("StartPage");return(0,a.uX)(),(0,a.Wv)(n)}var s=t(4232);const n=e=>((0,a.Qi)("data-v-00550b2a"),e=e(),(0,a.jt)(),e),c={class:"header"},i={key:0},u={key:1},d={class:"filter-zone"},f={key:0,class:"node-card"},h=n((()=>(0,a.Lk)("div",{class:"alert alert-success",role:"alert"}," Loading data ",-1))),p=[h],v={key:0,class:"tooltip-wrapper"},g={class:"graph-zone"},m=["font-size"],k=n((()=>(0,a.Lk)("br",null,null,-1))),b=n((()=>(0,a.Lk)("br",null,null,-1)));function w(e,o,t,l,n,h){const w=(0,a.g2)("ErrorNotification"),y=(0,a.g2)("FilterBar"),L=(0,a.g2)("NodeCard"),x=(0,a.g2)("VNetworkGraph");return(0,a.uX)(),(0,a.CE)(a.FK,null,[(0,a.bF)(w,{"error-message":e.errorMessage},null,8,["error-message"]),(0,a.Lk)("div",c,[e.isGraphVisible?((0,a.uX)(),(0,a.CE)("div",i,[(0,a.Lk)("div",null,[(0,a.Lk)("h5",null," Dependency graph for: "+(0,s.v_)(e.data.moduleName),1)])])):(0,a.Q3)("",!0),e.isGraphVisible?((0,a.uX)(),(0,a.CE)("div",u,[(0,a.Lk)("div",d,[(0,a.bF)(y,{filter:e.filter,onFilterChanged:e.onFilterChanged},null,8,["filter","onFilterChanged"])]),e.selectedNode?((0,a.uX)(),(0,a.CE)("div",f,[(0,a.bF)(L,{focusFilter:e.focusFilter,selectedNode:e.selectedNode,onFocusRequest:e.onFilterChanged},null,8,["focusFilter","selectedNode","onFocusRequest"])])):(0,a.Q3)("",!0)])):(0,a.Q3)("",!0),(0,a.bF)(r.eB,null,{default:(0,a.k6)((()=>[e.isLoaderVisible?((0,a.uX)(),(0,a.CE)("div",{key:0,class:(0,s.C4)(["centered-div loader",{"fade-enter":e.isLoaderVisible,"fade-leave-to":!e.isLoaderVisible}])},p,2)):(0,a.Q3)("",!0)])),_:1})]),e.isGraphVisible?((0,a.uX)(),(0,a.CE)("div",v,[(0,a.Lk)("div",g,[(0,a.bF)(x,{class:"graph",ref:"graph",layouts:e.layouts,"onUpdate:layouts":o[0]||(o[0]=o=>e.layouts=o),"selected-nodes":e.selectedNodes,"onUpdate:selectedNodes":o[1]||(o[1]=o=>e.selectedNodes=o),nodes:e.data.nodes,edges:e.data.edges,configs:e.configs,"event-handlers":e.eventHandlers},{"override-node-label":(0,a.k6)((({scale:e,text:o})=>[(0,a.Lk)("text",{x:"0",y:"0","font-size":14*e,"text-anchor":"middle","dominant-baseline":"central",fill:"#000000","font-family":"monospace"},(0,s.v_)(o),9,m)])),_:1},8,["layouts","selected-nodes","nodes","edges","configs","event-handlers"])]),(0,a.Lk)("div",{ref:"tooltip",class:"tooltip",style:(0,s.Tr)({...e.tooltipPos,opacity:e.tooltipOpacity})},[(0,a.Lk)("div",null,[(0,a.eW)((0,s.v_)(e.data.nodes[e.targetNodeId]?.name??"")+" ",1),k,(0,a.eW)(" FanIn: "+(0,s.v_)(e.data.nodes[e.targetNodeId]?.fanInCount??-1)+" ",1),b,(0,a.eW)(" FanOut: "+(0,s.v_)(e.data.nodes[e.targetNodeId]?.fanOutCount??-1),1)])],4),(0,a.Lk)("div",{ref:"edgetip",class:"edgetip",style:(0,s.Tr)({...e.edgetipPos,opacity:e.edgetipOpacity})},[(0,a.Lk)("div",null,(0,s.v_)(e.data.nodes[e.data.edges[e.targetEdgeId]?.source]?.name??" ??")+" -> "+(0,s.v_)(e.data.nodes[e.data.edges[e.targetEdgeId]?.target]?.name??" ??"),1)],4)])):(0,a.Q3)("",!0)],64)}var y=t(144),L=t(7266),x=t(8355);const F=e=>((0,a.Qi)("data-v-7e1fe430"),e=e(),(0,a.jt)(),e),O={class:"filter-container"},R={class:"form-check form-switch"},C=F((()=>(0,a.Lk)("label",{class:"form-check-label",for:"showStandard"},"Show standard",-1))),_={class:"form-check form-switch"},N=F((()=>(0,a.Lk)("label",{class:"form-check-label",for:"showPlatform"},"Show platform",-1))),S={class:"form-check form-switch"},E=F((()=>(0,a.Lk)("label",{class:"form-check-label",for:"showOuter"},"Show outer",-1))),K={class:"form-check form-switch"},I=F((()=>(0,a.Lk)("label",{class:"form-check-label",for:"showRatio"},"Show ratio",-1)));function P(e,o,t,l,s,n){return(0,a.uX)(),(0,a.CE)("div",O,[(0,a.Lk)("form",null,[(0,a.Lk)("div",R,[(0,a.bo)((0,a.Lk)("input",{class:"form-check-input",type:"checkbox","onUpdate:modelValue":o[0]||(o[0]=e=>n.filterLocal.showStandard=e),id:"showStandard"},null,512),[[r.lH,n.filterLocal.showStandard]]),C]),(0,a.Lk)("div",_,[(0,a.bo)((0,a.Lk)("input",{class:"form-check-input",type:"checkbox","onUpdate:modelValue":o[1]||(o[1]=e=>n.filterLocal.showPlatform=e),id:"showPlatform"},null,512),[[r.lH,n.filterLocal.showPlatform]]),N]),(0,a.Lk)("div",S,[(0,a.bo)((0,a.Lk)("input",{class:"form-check-input",type:"checkbox","onUpdate:modelValue":o[2]||(o[2]=e=>n.filterLocal.showOuter=e),id:"showOuter"},null,512),[[r.lH,n.filterLocal.showOuter]]),E]),(0,a.Lk)("div",K,[(0,a.bo)((0,a.Lk)("input",{class:"form-check-input",type:"checkbox","onUpdate:modelValue":o[3]||(o[3]=e=>n.filterLocal.showRatio=e),id:"showRatio"},null,512),[[r.lH,n.filterLocal.showRatio]]),I])])])}var V={name:"FilterBar",props:["filter"],computed:{filterLocal:{get:function(){return this.filter}}},created(){this.$watch("filterLocal",(e=>{this.$emit("filter-changed",e)}),{deep:!0})}},j=t(1241);const z=(0,j.A)(V,[["render",P],["__scopeId","data-v-7e1fe430"]]);var H=z;const M=e=>((0,a.Qi)("data-v-795db1b9"),e=e(),(0,a.jt)(),e),T={key:0},q={class:"header"},Q={class:"filter-container"},X={class:"form-check form-switch"},G=M((()=>(0,a.Lk)("label",{class:"form-check-label",for:"focusIns"},"Show ins",-1))),U={class:"form-check form-switch"},W=M((()=>(0,a.Lk)("label",{class:"form-check-label",for:"focusOuts"},"Show outs",-1)));function A(e,o,t,l,n,c){return t.selectedNode?((0,a.uX)(),(0,a.CE)("div",T,[(0,a.Lk)("div",q,(0,s.v_)(t.selectedNode.name),1),(0,a.Lk)("div",Q,[(0,a.Lk)("form",null,[(0,a.Lk)("div",null,[(0,a.Lk)("div",X,[(0,a.bo)((0,a.Lk)("input",{class:"form-check-input",type:"checkbox","onUpdate:modelValue":o[0]||(o[0]=e=>c.filterLocal.focusIns=e),id:"focusIns",required:""},null,512),[[r.lH,c.filterLocal.focusIns]]),G]),(0,a.Lk)("div",U,[(0,a.bo)((0,a.Lk)("input",{class:"form-check-input",type:"checkbox","onUpdate:modelValue":o[1]||(o[1]=e=>c.filterLocal.focusOuts=e),id:"focusOuts",required:""},null,512),[[r.lH,c.filterLocal.focusOuts]]),W])]),(0,a.Lk)("button",{type:"button",class:"btn btn-primary",onClick:o[2]||(o[2]=e=>c.fireFocusRequest(c.filterLocal,!0))},"Focus"),(0,a.Lk)("button",{type:"button",class:"btn btn-primary",onClick:o[3]||(o[3]=e=>c.fireFocusRequest(c.filterLocal,!1))},"Reset focus")])])])):(0,a.Q3)("",!0)}var B={name:"NodeCard",props:["focusFilter","selectedNode"],computed:{filterLocal:{get:function(){return this.focusFilter}}},methods:{fireFocusRequest(e,o){e.doFocus=o,this.$emit("focus-request",e)}}};const D=(0,j.A)(B,[["render",A],["__scopeId","data-v-795db1b9"]]);var $=D;const J={key:0},Y={class:"error-notification alert alert-danger"};function Z(e,o,t,r,l,n){return e.showMessage?((0,a.uX)(),(0,a.CE)("div",J,[(0,a.Lk)("div",Y,[(0,a.Lk)("button",{class:"btn btn-close top-right",onClick:o[0]||(o[0]=(...o)=>e.closeError&&e.closeError(...o))}),(0,a.Lk)("p",null,(0,s.v_)(e.$props.errorMessage),1)])])):(0,a.Q3)("",!0)}var ee={name:"ErrorNotification",props:{errorMessage:{required:!0}},watch:{errorMessage(e){e.length>0&&this.showError()}},setup:function(){const e=(0,y.KR)(!1),o=function(){e.value=!1},t=(0,y.KR)(5e3),r=function(){e.value=!0,setTimeout((()=>{o()}),t.value)};return{showMessage:e,closeError:o,showError:r,timeout:t}}};const oe=(0,j.A)(ee,[["render",Z],["__scopeId","data-v-0573c0fc"]]);var te=oe,re={name:"StartPage",components:{FilterBar:H,VNetworkGraph:L.zS,NodeCard:$,ErrorNotification:te},mounted(){this.showGraph()},setup:function(){const e=(0,y.KR)(!1),o=(0,y.KR)(!1),t=(0,y.KR)([]),r=(0,y.KR)(),l=(0,y.KR)(""),s=(0,y.Kh)({nodes:{},edges:{},moduleName:"loading"}),n=(0,y.KR)({}),c=()=>{e.value=!0,u()},i=function(){c()},u=async()=>{try{let e=0;H.focusIns&&H.focusOuts?e=2:H.focusIns&&(e=1);let t="";H.doFocus&&(t=r.value?.id??"");const a=await x.A.get("/api/graph",{params:{show_standard:z.showStandard,show_platform:z.showPlatform,show_outer:z.showOuter,focus_package:t,focus_type:e,show_ratio:z.showRatio}});s.nodes=a.data.nodes,s.edges=a.data.edges,s.moduleName=a.data.moduleName,o.value=!0}catch(t){d("Cannot fetch graph: "+t.message)}e.value=!1},d=function(e){const o=(new Date).toLocaleTimeString("ru-ru",{hour12:!1});l.value=o+" "+e},f=20,h=f/2,p=8,v=25,g=10,m=function(e){let o=(e.name.length-1)*p+v;return z.showRatio||(o+=(e.fanInCount.toString().length+e.fanOutCount.toString().length+2)*p),Math.ceil(o/g)*g},k=L.xq({view:{layoutHandler:new L.JT({grid:g})},node:{selectable:!0,normal:{radius:e=>e.size??h,color:e=>e.color,type:"rect",width:e=>m(e)},hover:{color:e=>e.color},label:{color:"#000000",fontSize:14}},edge:{type:"curve",normal:{color:e=>e.color,width:2},hover:{color:e=>e.color},margin:4,marker:{target:{type:"arrow",width:4,height:4}}}}),b=(0,y.KR)(),w=(0,y.KR)(),F=(0,y.KR)(""),O=(0,y.KR)(0),R=(0,y.KR)({left:"0px",top:"0px"}),C=(0,y.KR)(),_=(0,y.KR)(0),N=(0,y.KR)({left:"0px",top:"0px"}),S=2,E=(0,y.KR)(""),K=(0,a.EW)((()=>{if(!F.value)return{x:0,y:0};const e=n.value.nodes[F.value];return e||{x:0,y:0}})),I=(0,a.EW)((()=>{const e=s.edges[E.value];if(!e)return{x:0,y:0};const o=s.edges[E.value].source,t=s.edges[E.value].target;return{x:(n.value.nodes[o].x+n.value.nodes[t].x)/2,y:(n.value.nodes[o].y+n.value.nodes[t].y)/2}}));(0,a.wB)((()=>[K.value,O.value]),(()=>{if(!b.value||!w.value)return;const e=b.value.translateFromSvgToDomCoordinates(K.value);R.value={left:e.x-w.value.offsetWidth/2+"px",top:e.y-h-w.value.offsetHeight-10+"px"}}),{deep:!0}),(0,a.wB)((()=>[I.value,_.value]),(()=>{if(!b.value||!C.value)return{x:0,y:0};if(!E.value)return{x:0,y:0};const e=b.value.translateFromSvgToDomCoordinates(I.value);N.value={left:e.x-C.value.offsetWidth/2+"px",top:e.y-S-C.value.offsetHeight-10+"px"}}),{deep:!0}),(0,a.wB)((()=>[t.value]),(()=>{0===t.value.length?r.value=null:r.value=s.nodes[t.value[0]]}),{deep:!0});const P=(0,y.KR)(""),V=(0,y.KR)(""),j={"node:pointerover":({node:e})=>{F.value=e,O.value=1},"node:pointerout":({node:e})=>{F.value=e,O.value=0},"edge:pointerover":({edge:e})=>{E.value=e??"";const o=s.edges[e];P.value=s.nodes[o.source].color,s.nodes[o.source].color="red",s.nodes[o.source].size=2*h,V.value=s.nodes[o.target].color,s.nodes[o.target].color="red",s.nodes[o.target].size=2*h,_.value=1},"edge:pointerout":({edge:e})=>{E.value=e??"",_.value=0;const o=s.edges[e];s.nodes[o.source].color=P.value,delete s.nodes[o.source].size,s.nodes[o.target].color=V.value,delete s.nodes[o.target].size}},z=(0,y.Kh)({showStandard:!1,showPlatform:!0,showOuter:!0,showRatio:!1}),H=(0,y.Kh)({focusIns:!0,focusOuts:!0,doFocus:!1});return{isLoaderVisible:e,isGraphVisible:o,showGraph:c,data:s,layouts:n,configs:k,eventHandlers:j,tooltipPos:R,tooltipOpacity:O,targetNodeId:F,graph:b,tooltip:w,edgetip:C,targetEdgeId:E,edgetipPos:N,edgetipOpacity:_,edgeCenterPos:I,filter:z,onFilterChanged:i,selectedNodes:t,selectedNode:r,focusFilter:H,errorMessage:l}}};const ae=(0,j.A)(re,[["render",w],["__scopeId","data-v-00550b2a"]]);var le=ae,se={name:"App",components:{StartPage:le}};const ne=(0,j.A)(se,[["render",l]]);var ce=ne;(0,r.Ef)(ce).mount("#app")}},o={};function t(r){var a=o[r];if(void 0!==a)return a.exports;var l=o[r]={exports:{}};return e[r].call(l.exports,l,l.exports,t),l.exports}t.m=e,function(){var e=[];t.O=function(o,r,a,l){if(!r){var s=1/0;for(u=0;u<e.length;u++){r=e[u][0],a=e[u][1],l=e[u][2];for(var n=!0,c=0;c<r.length;c++)(!1&l||s>=l)&&Object.keys(t.O).every((function(e){return t.O[e](r[c])}))?r.splice(c--,1):(n=!1,l<s&&(s=l));if(n){e.splice(u--,1);var i=a();void 0!==i&&(o=i)}}return o}l=l||0;for(var u=e.length;u>0&&e[u-1][2]>l;u--)e[u]=e[u-1];e[u]=[r,a,l]}}(),function(){t.n=function(e){var o=e&&e.__esModule?function(){return e["default"]}:function(){return e};return t.d(o,{a:o}),o}}(),function(){t.d=function(e,o){for(var r in o)t.o(o,r)&&!t.o(e,r)&&Object.defineProperty(e,r,{enumerable:!0,get:o[r]})}}(),function(){t.g=function(){if("object"===typeof globalThis)return globalThis;try{return this||new Function("return this")()}catch(e){if("object"===typeof window)return window}}()}(),function(){t.o=function(e,o){return Object.prototype.hasOwnProperty.call(e,o)}}(),function(){t.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})}}(),function(){var e={524:0};t.O.j=function(o){return 0===e[o]};var o=function(o,r){var a,l,s=r[0],n=r[1],c=r[2],i=0;if(s.some((function(o){return 0!==e[o]}))){for(a in n)t.o(n,a)&&(t.m[a]=n[a]);if(c)var u=c(t)}for(o&&o(r);i<s.length;i++)l=s[i],t.o(e,l)&&e[l]&&e[l][0](),e[l]=0;return t.O(u)},r=self["webpackChunkvue_app"]=self["webpackChunkvue_app"]||[];r.forEach(o.bind(null,0)),r.push=o.bind(null,r.push.bind(r))}();var r=t.O(void 0,[504],(function(){return t(1627)}));r=t.O(r)})();
//# sourceMappingURL=app.c40083e5.js.map