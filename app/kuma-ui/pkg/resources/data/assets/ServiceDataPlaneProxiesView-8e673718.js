import{E as S}from"./ErrorBlock-54d14d19.js";import{D as C,F as V}from"./FilterBar-95a88cae.js";import{S as x}from"./SummaryView-f4c41312.js";import{d as z,a as t,o as i,b as l,w as s,e as r,m as P,f as n,t as k,B as q,p as u,_ as T}from"./index-36b38e0c.js";import"./index-fce48c05.js";import"./TextWithCopyButton-69a2e47a.js";import"./CopyButton-dbd4bffe.js";import"./WarningIcon.vue_vue_type_script_setup_true_lang-7d939fae.js";import"./AppCollection-bc90608e.js";import"./StatusBadge-111d6c49.js";import"./uniqueId-90cc9b93.js";const B=z({__name:"ServiceDataPlaneProxiesView",setup(R){return($,N)=>{const y=t("RouteTitle"),f=t("KSelect"),g=t("KCard"),v=t("RouterView"),m=t("DataSource"),w=t("AppView"),h=t("RouteView");return i(),l(m,{src:"/me"},{default:s(({data:c})=>[c?(i(),l(h,{key:0,name:"service-data-plane-proxies-view",params:{page:1,size:c.pageSize,query:"",dataplaneType:"all",s:"",mesh:"",service:"",dataPlane:""}},{default:s(({can:b,route:e,t:d})=>[r(w,null,{title:s(()=>[P("h2",null,[r(y,{title:d("services.routes.item.navigation.service-data-plane-proxies-view")},null,8,["title"])])]),default:s(()=>[n(),r(m,{src:`/meshes/${e.params.mesh}/dataplanes/for/${e.params.service}/of/${e.params.dataplaneType}?page=${e.params.page}&size=${e.params.size}&search=${e.params.s}`},{default:s(({data:o,error:p})=>[r(g,null,{default:s(()=>[p!==void 0?(i(),l(S,{key:0,error:p},null,8,["error"])):(i(),l(C,{key:1,"data-testid":"data-plane-collection","page-number":e.params.page,"page-size":e.params.size,total:o==null?void 0:o.total,items:o==null?void 0:o.items,error:p,"is-selected-row":a=>a.name===e.params.dataPlane,"summary-route-name":"service-data-plane-summary-view","is-global-mode":b("use zones"),onChange:e.update},{toolbar:s(()=>[r(V,{class:"data-plane-proxy-filter",placeholder:"tag: 'kuma.io/protocol: http'",query:e.params.query,fields:{name:{description:"filter by name or parts of a name"},protocol:{description:"filter by “kuma.io/protocol” value"},tag:{description:"filter by tags (e.g. “tag: version:2”)"},zone:{description:"filter by “kuma.io/zone” value"}},onFieldsChange:a=>e.update({query:a.query,s:a.query.length>0?JSON.stringify(a.fields):""})},null,8,["placeholder","query","fields","onFieldsChange"]),n(),r(f,{class:"filter-select",label:"Type",items:["all","standard","builtin","delegated"].map(a=>({value:a,label:d(`data-planes.type.${a}`),selected:a===e.params.dataplaneType})),onSelected:a=>e.update({dataplaneType:String(a.value)})},{"item-template":s(({item:a})=>[n(k(a.label),1)]),_:2},1032,["items","onSelected"])]),_:2},1032,["page-number","page-size","total","items","error","is-selected-row","is-global-mode","onChange"]))]),_:2},1024),n(),e.params.dataPlane?(i(),l(v,{key:0},{default:s(a=>[r(x,{onClose:_=>e.replace({name:"service-data-plane-proxies-view",params:{mesh:e.params.mesh},query:{page:e.params.page,size:e.params.size}})},{default:s(()=>[(i(),l(q(a.Component),{name:e.params.dataPlane,"dataplane-overview":o==null?void 0:o.items.find(_=>_.name===e.params.dataPlane)},null,8,["name","dataplane-overview"]))]),_:2},1032,["onClose"])]),_:2},1024)):u("",!0)]),_:2},1032,["src"])]),_:2},1024)]),_:2},1032,["params"])):u("",!0)]),_:1})}}});const M=T(B,[["__scopeId","data-v-daea0c6b"]]);export{M as default};