import{d as B,r as l,o as s,m as r,w as t,b as n,k as b,A as $,e as i,t as m,c as p,F as u,S as E,p as d,E as P,q as I}from"./index-BRBWbknO.js";import{F as L}from"./FilterBar-B_LXsRot.js";import{S as N}from"./SummaryView-CK1e_JHQ.js";const X={class:"stack"},q={key:0},T={key:1},F=B({__name:"BuiltinGatewayDataplanesView",setup(R){return(G,K)=>{const _=l("XAction"),v=l("XIcon"),w=l("XActionGroup"),z=l("RouterView"),C=l("DataLoader"),x=l("KCard"),f=l("DataSource"),V=l("AppView"),S=l("RouteView");return s(),r(f,{src:"/me"},{default:t(({data:h})=>[h?(s(),r(S,{key:0,name:"builtin-gateway-dataplanes-view",params:{mesh:"",gateway:"",listener:"",page:1,size:h.pageSize,s:"",dataPlane:""}},{default:t(({can:k,route:a,t:o})=>[n(V,null,{default:t(()=>[n(f,{src:`/meshes/${a.params.mesh}/mesh-gateways/${a.params.gateway}`},{default:t(({data:y,error:A})=>[b("div",X,[n(x,null,{default:t(()=>[n(C,{src:y===void 0?"":`/meshes/${a.params.mesh}/dataplanes/for/service-insight/${y.selectors[0].match["kuma.io/service"]}?page=${a.params.page}&size=${a.params.size}&search=${a.params.s}`,data:[y],errors:[A],loader:!1},{default:t(({data:c})=>[n($,{class:"data-plane-collection","data-testid":"data-plane-collection","page-number":a.params.page,"page-size":a.params.size,headers:[{label:"Name",key:"name"},{label:"Namespace",key:"namespace"},...k("use zones")?[{label:"Zone",key:"zone"}]:[],{label:"Certificate Info",key:"certificate"},{label:"Status",key:"status"},{label:"Warnings",key:"warnings",hideLabel:!0},{label:"Actions",key:"actions",hideLabel:!0}],items:c==null?void 0:c.items,total:c==null?void 0:c.total,"is-selected-row":e=>e.name===a.params.dataPlane,"summary-route-name":"builtin-gateway-data-plane-summary-view","empty-state-message":o("common.emptyState.message",{type:"Data Plane Proxies"}),"empty-state-cta-to":o("data-planes.href.docs.data_plane_proxy"),"empty-state-cta-text":o("common.documentation"),onChange:a.update},{toolbar:t(()=>[n(L,{class:"data-plane-proxy-filter",placeholder:"name:dataplane-name",query:a.params.s,fields:{name:{description:"filter by name or parts of a name"},protocol:{description:"filter by “kuma.io/protocol” value"},service:{description:"filter by “kuma.io/service” value"},tag:{description:"filter by tags (e.g. “tag: version:2”)"},...k("use zones")&&{zone:{description:"filter by “kuma.io/zone” value"}}},onChange:e=>a.update({...Object.fromEntries(e.entries())})},null,8,["query","fields","onChange"])]),namespace:t(({row:e})=>[i(m(e.namespace),1)]),name:t(({row:e})=>[n(_,{"data-action":"",class:"name-link",title:e.name,to:{name:"builtin-gateway-data-plane-summary-view",params:{mesh:e.mesh,dataPlane:e.id},query:{page:a.params.page,size:a.params.size,s:a.params.s}}},{default:t(()=>[i(m(e.name),1)]),_:2},1032,["title","to"])]),zone:t(({row:e})=>[e.zone?(s(),r(_,{key:0,to:{name:"zone-cp-detail-view",params:{zone:e.zone}}},{default:t(()=>[i(m(e.zone),1)]),_:2},1032,["to"])):(s(),p(u,{key:1},[i(m(o("common.collection.none")),1)],64))]),certificate:t(({row:e})=>{var g;return[(g=e.dataplaneInsight.mTLS)!=null&&g.certificateExpirationTime?(s(),p(u,{key:0},[i(m(o("common.formats.datetime",{value:Date.parse(e.dataplaneInsight.mTLS.certificateExpirationTime)})),1)],64)):(s(),p(u,{key:1},[i(m(o("data-planes.components.data-plane-list.certificate.none")),1)],64))]}),status:t(({row:e})=>[n(E,{status:e.status},null,8,["status"])]),warnings:t(({row:e})=>[e.isCertExpired||e.warnings.length>0?(s(),r(v,{key:0,class:"mr-1",name:"warning"},{default:t(()=>[b("ul",null,[e.warnings.length>0?(s(),p("li",q,m(o("data-planes.components.data-plane-list.version_mismatch")),1)):d("",!0),i(),e.isCertExpired?(s(),p("li",T,m(o("data-planes.components.data-plane-list.cert_expired")),1)):d("",!0)])]),_:2},1024)):(s(),p(u,{key:1},[i(m(o("common.collection.none")),1)],64))]),actions:t(({row:e})=>[n(w,null,{default:t(()=>[n(_,{to:{name:"data-plane-detail-view",params:{dataPlane:e.id}}},{default:t(()=>[i(m(o("common.collection.actions.view")),1)]),_:2},1032,["to"])]),_:2},1024)]),_:2},1032,["page-number","page-size","headers","items","total","is-selected-row","empty-state-message","empty-state-cta-to","empty-state-cta-text","onChange"]),i(),a.params.dataPlane?(s(),r(z,{key:0},{default:t(e=>[n(N,{onClose:g=>a.replace({name:a.name,params:{mesh:a.params.mesh},query:{page:a.params.page,size:a.params.size,s:a.params.s}})},{default:t(()=>[typeof c<"u"?(s(),r(P(e.Component),{key:0,items:c.items},null,8,["items"])):d("",!0)]),_:2},1032,["onClose"])]),_:2},1024)):d("",!0)]),_:2},1032,["src","data","errors"])]),_:2},1024)])]),_:2},1032,["src"])]),_:2},1024)]),_:2},1032,["params"])):d("",!0)]),_:1})}}}),Z=I(F,[["__scopeId","data-v-990d8133"]]);export{Z as default};