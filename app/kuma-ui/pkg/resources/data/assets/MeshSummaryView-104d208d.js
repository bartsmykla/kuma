import{d as v,k as w,a as c,o as _,b as d,w as e,e as l,m as s,f as t,l as n,O as f,t as m,c as g,R as p,s as k,v as I,U as V,_ as x}from"./index-3ddd0e9e.js";const R=i=>(k("data-v-5ee23800"),i=i(),I(),i),S={class:"summary-title-wrapper"},B=R(()=>s("img",{"aria-hidden":"true",src:V},null,-1)),M={class:"summary-title"},T={key:1,class:"stack"},C={class:"mt-4 stack"},N=v({__name:"MeshSummaryView",props:{name:{},meshInsight:{default:void 0}},setup(i){const{t:a}=w(),o=i;return(A,L)=>{const r=c("RouteTitle"),u=c("RouterLink"),h=c("AppView"),y=c("RouteView");return _(),d(y,{name:"mesh-summary-view"},{default:e(()=>[l(h,null,{title:e(()=>[s("div",S,[B,t(),s("h2",M,[l(u,{to:{name:"mesh-detail-view",params:{mesh:o.name}}},{default:e(()=>[l(r,{title:n(a)("meshes.routes.item.title",{name:o.name})},null,8,["title"])]),_:1},8,["to"])])])]),default:e(()=>[t(),o.meshInsight===void 0?(_(),d(f,{key:0},{message:e(()=>[s("p",null,m(n(a)("common.collection.summary.empty_message",{type:"Mesh"})),1)]),default:e(()=>[t(m(n(a)("common.collection.summary.empty_title",{type:"Mesh"}))+" ",1)]),_:1})):(_(),g("div",T,[s("div",null,[s("h3",null,m(n(a)("meshes.routes.item.overview")),1),t(),s("div",C,[l(p,{total:o.meshInsight.services.total,"data-testid":"services-status"},{title:e(()=>[t(m(n(a)("meshes.detail.services")),1)]),_:1},8,["total"]),t(),l(p,{online:o.meshInsight.dataplanesByType.standard.online,total:o.meshInsight.dataplanesByType.standard.total,"data-testid":"data-plane-proxies-status"},{title:e(()=>[t(m(n(a)("meshes.detail.data_plane_proxies")),1)]),_:1},8,["online","total"])])])]))]),_:1})]),_:1})}}});const D=x(N,[["__scopeId","data-v-5ee23800"]]);export{D as default};
