import{d as v,k as g,o as c,c as _,e as n,w as e,f as t,t as m,l as a,Y as u,b as d,F as k,a as l,m as i,O as x,s as S,v as V,aE as O,_ as B}from"./index-671fe4fd.js";import{S as R}from"./StatusBadge-d7b39d12.js";import{T as Z}from"./TextWithCopyButton-fcf481cf.js";import"./CopyButton-2a28adb7.js";import"./index-fce48c05.js";const C={class:"stack"},I=v({__name:"ZoneEgressSummary",props:{zoneEgressOverview:{}},setup(o){const{t:s}=g(),r=o;return(y,w)=>(c(),_("div",C,[n(u,null,{title:e(()=>[t(m(a(s)("http.api.property.status")),1)]),body:e(()=>[n(R,{status:r.zoneEgressOverview.state},null,8,["status"])]),_:1}),t(),n(u,null,{title:e(()=>[t(m(a(s)("http.api.property.address")),1)]),body:e(()=>[r.zoneEgressOverview.zoneEgress.socketAddress.length>0?(c(),d(Z,{key:0,text:r.zoneEgressOverview.zoneEgress.socketAddress},null,8,["text"])):(c(),_(k,{key:1},[t(m(a(s)("common.detail.none")),1)],64))]),_:1})]))}}),T=o=>(S("data-v-01cd37d3"),o=o(),V(),o),A={class:"summary-title-wrapper"},b=T(()=>i("img",{"aria-hidden":"true",src:O},null,-1)),N={class:"summary-title"},$={key:1,class:"stack"},D=v({__name:"ZoneEgressSummaryView",props:{zoneEgressOverview:{default:void 0}},setup(o){const{t:s}=g(),r=o;return(y,w)=>{const z=l("RouteTitle"),E=l("RouterLink"),f=l("AppView"),h=l("RouteView");return c(),d(h,{name:"zone-egress-summary-view",params:{zone:"",zoneEgress:""}},{default:e(({route:p})=>[n(f,null,{title:e(()=>[i("div",A,[b,t(),i("h2",N,[n(E,{to:{name:"zone-egress-detail-view",params:{zone:p.params.zone,zoneEgress:p.params.zoneEgress}}},{default:e(()=>[n(z,{title:a(s)("zone-egresses.routes.item.title",{name:p.params.zoneEgress})},null,8,["title"])]),_:2},1032,["to"])])])]),default:e(()=>[t(),r.zoneEgressOverview===void 0?(c(),d(x,{key:0},{message:e(()=>[i("p",null,m(a(s)("common.collection.summary.empty_message",{type:"ZoneEgress"})),1)]),default:e(()=>[t(m(a(s)("common.collection.summary.empty_title",{type:"ZoneEgress"}))+" ",1)]),_:1})):(c(),_("div",$,[i("div",null,[i("h3",null,m(a(s)("zone-egresses.routes.item.overview")),1),t(),n(I,{class:"mt-4","zone-egress-overview":r.zoneEgressOverview},null,8,["zone-egress-overview"])])]))]),_:2},1024)]),_:1})}}});const q=B(D,[["__scopeId","data-v-01cd37d3"]]);export{q as default};