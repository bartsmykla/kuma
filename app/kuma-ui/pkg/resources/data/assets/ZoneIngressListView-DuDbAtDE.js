import{d as S,a as i,o as t,b as a,w as s,e as o,m as V,f as m,a3 as x,t as c,T as g,c as k,F as y,p as z,P as h,K as b,E as R,q as p,_ as B}from"./index-CImj3nNu.js";import{A as N}from"./AppCollection-gXdGV6dw.js";import{S as D}from"./StatusBadge-P5eqtkea.js";import{S as L}from"./SummaryView-8V3W9fEN.js";const T=S({__name:"ZoneIngressListView",setup(E){return(K,Z)=>{const f=i("RouteTitle"),u=i("RouterLink"),w=i("KCard"),I=i("RouterView"),_=i("DataSource"),v=i("AppView"),C=i("RouteView");return t(),a(_,{src:"/me"},{default:s(({data:A})=>[A?(t(),a(C,{key:0,name:"zone-ingress-list-view",params:{zone:"",zoneIngress:""}},{default:s(({route:r,t:l})=>[o(v,null,{title:s(()=>[V("h2",null,[o(f,{title:l("zone-ingresses.routes.items.title")},null,8,["title"])])]),default:s(()=>[m(),o(_,{src:`/zone-cps/${r.params.zone}/ingresses?page=1&size=100`},{default:s(({data:n,error:d})=>[o(w,null,{default:s(()=>[d!==void 0?(t(),a(x,{key:0,error:d},null,8,["error"])):(t(),a(N,{key:1,class:"zone-ingress-collection","data-testid":"zone-ingress-collection",headers:[{label:"Name",key:"name"},{label:"Address",key:"socketAddress"},{label:"Advertised address",key:"advertisedSocketAddress"},{label:"Status",key:"status"},{label:"Details",key:"details",hideLabel:!0}],"page-number":1,"page-size":100,total:n==null?void 0:n.total,items:n==null?void 0:n.items,error:d,"empty-state-message":l("common.emptyState.message",{type:"Zone Ingresses"}),"empty-state-cta-to":l("zone-ingresses.href.docs"),"empty-state-cta-text":l("common.documentation"),"is-selected-row":e=>e.name===r.params.zoneIngress,onChange:r.update},{name:s(({row:e})=>[o(u,{to:{name:"zone-ingress-summary-view",params:{zone:r.params.zone,zoneIngress:e.id},query:{page:1,size:100}}},{default:s(()=>[m(c(e.name),1)]),_:2},1032,["to"])]),socketAddress:s(({row:e})=>[e.zoneIngress.socketAddress.length>0?(t(),a(g,{key:0,text:e.zoneIngress.socketAddress},null,8,["text"])):(t(),k(y,{key:1},[m(c(l("common.collection.none")),1)],64))]),advertisedSocketAddress:s(({row:e})=>[e.zoneIngress.advertisedSocketAddress.length>0?(t(),a(g,{key:0,text:e.zoneIngress.advertisedSocketAddress},null,8,["text"])):(t(),k(y,{key:1},[m(c(l("common.collection.none")),1)],64))]),status:s(({row:e})=>[o(D,{status:e.state},null,8,["status"])]),details:s(({row:e})=>[o(u,{class:"details-link","data-testid":"details-link",to:{name:"zone-ingress-detail-view",params:{zoneIngress:e.id}}},{default:s(()=>[m(c(l("common.collection.details_link"))+" ",1),o(z(h),{decorative:"",size:z(b)},null,8,["size"])]),_:2},1032,["to"])]),_:2},1032,["total","items","error","empty-state-message","empty-state-cta-to","empty-state-cta-text","is-selected-row","onChange"]))]),_:2},1024),m(),r.params.zoneIngress?(t(),a(I,{key:0},{default:s(e=>[o(L,{onClose:$=>r.replace({name:"zone-ingress-list-view",params:{zone:r.params.zone},query:{page:1,size:100}})},{default:s(()=>[typeof n<"u"?(t(),a(R(e.Component),{key:0,items:n.items},null,8,["items"])):p("",!0)]),_:2},1032,["onClose"])]),_:2},1024)):p("",!0)]),_:2},1032,["src"])]),_:2},1024)]),_:1},8,["params"])):p("",!0)]),_:1})}}}),U=B(T,[["__scopeId","data-v-ead0e5f6"]]);export{U as default};
