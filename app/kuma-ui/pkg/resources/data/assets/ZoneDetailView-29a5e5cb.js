import{d as x,c as v,r as E,o as t,a,w as e,q as V,g as d,h as r,s as z,e as k,F as O,v as I,b as _,R as N,X as T,u as $,j as A}from"./index-8624ce1f.js";import{a as L,A as Z,_ as R,S as j}from"./SubscriptionHeader.vue_vue_type_script_setup_true_lang-7224c3ac.js";import{_ as q}from"./CodeBlock.vue_vue_type_style_index_0_lang-80823b70.js";import{D as F,a as M}from"./DefinitionListItem-3a98eb59.js";import{T as P}from"./TabsWidget-a09e9d7e.js";import{T as S}from"./TextWithCopyButton-5d51ccad.js";import{_ as W}from"./WarningsWidget.vue_vue_type_script_setup_true_lang-c123b101.js";import{g as B,z as G,o as J,B as K,E as X,n as H,A as U,h as Q,_ as Y}from"./RouteView.vue_vue_type_script_setup_true_lang-d02ca438.js";import{_ as ee}from"./RouteTitle.vue_vue_type_script_setup_true_lang-73a33c91.js";import{_ as ne}from"./EmptyBlock.vue_vue_type_script_setup_true_lang-09d80466.js";import{E as te}from"./ErrorBlock-6921ebf1.js";const se={class:"entity-heading"},ae=x({__name:"ZoneDetails",props:{zoneOverview:{type:Object,required:!0}},setup(C){const i=C,{t:h}=B(),w=G(),p=[{hash:"#overview",title:"Overview"},{hash:"#insights",title:"Zone Insights"},{hash:"#config",title:"Config"},{hash:"#warnings",title:"Warnings"}],f=v(()=>({name:"zone-cp-detail-view",params:{zone:i.zoneOverview.name}})),c=v(()=>{const{type:n,name:s}=i.zoneOverview,l=J(i.zoneOverview.zoneInsight);return{type:n,name:s,status:l,"Authentication Type":K(i.zoneOverview)}}),y=v(()=>{var s;const n=((s=i.zoneOverview.zoneInsight)==null?void 0:s.subscriptions)??[];return Array.from(n).reverse()}),b=v(()=>{var l;const n=[],s=((l=i.zoneOverview.zoneInsight)==null?void 0:l.subscriptions)??[];if(s.length>0){const o=s[s.length-1],u=o.version.kumaCp.version||"-",{kumaCpGlobalCompatible:D=!0}=o.version.kumaCp;D||n.push({kind:X,payload:{zoneCpVersion:u,globalCpVersion:w("KUMA_VERSION")}})}return n}),g=v(()=>{var s;const n=((s=i.zoneOverview.zoneInsight)==null?void 0:s.subscriptions)??[];if(n.length>0){const l=n[n.length-1];if(l.config)return JSON.stringify(JSON.parse(l.config),null,2)}return null}),m=v(()=>b.value.length===0?p.filter(n=>n.hash!=="#warnings"):p);return(n,s)=>{const l=E("router-link");return t(),a(P,{tabs:m.value},{tabHeader:e(()=>[V("h1",se,[d(`
        Zone Control Plane:

        `),r(S,{text:c.value.name},{default:e(()=>[r(l,{to:f.value},{default:e(()=>[d(z(c.value.name),1)]),_:1},8,["to"])]),_:1},8,["text"])])]),overview:e(()=>[r(F,null,{default:e(()=>[(t(!0),k(O,null,I(c.value,(o,u)=>(t(),a(M,{key:u,term:_(h)(`http.api.property.${u}`)},{default:e(()=>[u==="status"?(t(),a(_(N),{key:0,appearance:o==="offline"?"danger":"success"},{default:e(()=>[d(z(o),1)]),_:2},1032,["appearance"])):u==="name"?(t(),a(S,{key:1,text:o},null,8,["text"])):(t(),k(O,{key:2},[d(z(o),1)],64))]),_:2},1032,["term"]))),128))]),_:1})]),insights:e(()=>[r(L,{"initially-open":0},{default:e(()=>[(t(!0),k(O,null,I(y.value,(o,u)=>(t(),a(Z,{key:u},{"accordion-header":e(()=>[r(R,{details:o},null,8,["details"])]),"accordion-content":e(()=>[r(j,{details:o},null,8,["details"])]),_:2},1024))),128))]),_:1})]),config:e(()=>[g.value!==null?(t(),a(q,{key:0,id:"code-block-zone-config",language:"json",code:g.value,"is-searchable":"","query-key":"zone-config"},null,8,["code"])):(t(),a(_(T),{key:1,"data-testid":"warning-no-subscriptions",appearance:"warning"},{alertMessage:e(()=>[d(z(_(h)("zone-cps.routes.item.config.no-subscriptions")),1)]),_:1}))]),warnings:e(()=>[r(W,{warnings:b.value},null,8,["warnings"])]),_:1},8,["tabs"])}}}),oe={class:"zone-details"},re={key:3,class:"kcard-border","data-testid":"detail-view-details"},he=x({__name:"ZoneDetailView",setup(C){const i=H(),h=$(),{t:w}=B(),p=A(null),f=A(!0),c=A(null);y();function y(){b()}async function b(){f.value=!0,c.value=null;const g=h.params.zone;try{p.value=await i.getZoneOverview({name:g})}catch(m){p.value=null,m instanceof Error?c.value=m:console.error(m)}finally{f.value=!1}}return(g,m)=>(t(),a(Y,null,{default:e(({route:n})=>[r(ee,{title:_(w)("zone-cps.routes.item.title",{name:n.params.zone})},null,8,["title"]),d(),r(U,{breadcrumbs:[{to:{name:"zone-cp-list-view"},text:_(w)("zone-cps.routes.item.breadcrumbs")}]},{default:e(()=>[V("div",oe,[f.value?(t(),a(Q,{key:0})):c.value!==null?(t(),a(te,{key:1,error:c.value},null,8,["error"])):p.value===null?(t(),a(ne,{key:2})):(t(),k("div",re,[r(ae,{"zone-overview":p.value},null,8,["zone-overview"])]))])]),_:1},8,["breadcrumbs"])]),_:1}))}});export{he as default};