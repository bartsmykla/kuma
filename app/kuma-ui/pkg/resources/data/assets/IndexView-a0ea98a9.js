import{d as N,g as R,R as T,a4 as E,y as A,o as r,l as B,j as c,w as n,k as e,a3 as D,aM as M,n as _,H as C,a1 as I,i as d,p as g,m as Z,r as b,E as $,x as L,a8 as j,a5 as P,q}from"./index-23176b1b.js";import{_ as G}from"./DeleteResourceModal.vue_vue_type_script_setup_true_lang-471dba5f.js";import{N as Y}from"./NavTabs-4ef57897.js";const H=N({__name:"ZoneActionMenu",props:{zoneOverview:{type:Object,required:!0},kpopAttributes:{type:Object,default:()=>({placement:"bottomEnd"})}},setup(h){const{t}=R(),x=T(),O=E(),l=h,u=A(!1);function v(){u.value=!u.value}async function w(){await x.deleteZone({name:l.zoneOverview.name})}function o(){O.push({name:"zone-cp-list-view"})}return(i,a)=>(r(),B("div",null,[c(e(I),{"button-appearance":"creation","kpop-attributes":l.kpopAttributes,label:e(t)("zones.action_menu.toggle_button"),"show-caret":"",width:"280"},{items:n(()=>[c(e(D),{"is-dangerous":"","data-testid":"delete-button",onClick:M(v,["prevent"])},{default:n(()=>[_(C(e(t)("zones.action_menu.delete_button")),1)]),_:1},8,["onClick"])]),_:1},8,["kpop-attributes","label"]),_(),u.value?(r(),d(G,{key:0,"confirmation-text":l.zoneOverview.name,"delete-function":w,"is-visible":"","action-button-text":e(t)("common.delete_modal.proceed_button"),title:e(t)("common.delete_modal.title",{type:"Zone"}),"data-testid":"delete-zone-modal",onCancel:v,onDelete:o},{"body-content":n(()=>[g("p",null,C(e(t)("common.delete_modal.text1",{type:"Zone",name:l.zoneOverview.name})),1),_(),g("p",null,C(e(t)("common.delete_modal.text2")),1)]),_:1},8,["confirmation-text","action-button-text","title"])):Z("",!0)]))}}),F=N({__name:"IndexView",setup(h){var w;const{t}=R(),l=(((w=E().getRoutes().find(o=>o.name==="zone-cp-detail-tabs-view"))==null?void 0:w.children)??[]).map(o=>{var s,p;const i=typeof o.name>"u"?(s=o.children)==null?void 0:s[0]:o,a=i.name,m=((p=i.meta)==null?void 0:p.module)??"";return{title:t(`zone-cps.routes.item.navigation.${a}`),routeName:a,module:m}}),u=A([]),v=o=>{var m,f;const i=[],a=((m=o.zoneInsight)==null?void 0:m.subscriptions)??[];if(a.length>0){const s=a[a.length-1],p=s.version.kumaCp.version||"-",{kumaCpGlobalCompatible:z=!0}=s.version.kumaCp;s.config&&((f=JSON.parse(s.config))==null?void 0:f.store.type)==="memory"&&i.push({kind:"ZONE_STORE_TYPE_MEMORY",payload:{}}),z||i.push({kind:"INCOMPATIBLE_ZONE_AND_GLOBAL_CPS_VERSIONS",payload:{zoneCpVersion:p}})}u.value=i};return(o,i)=>{const a=b("RouteTitle"),m=b("RouterView"),f=b("AppView"),s=b("DataSource"),p=b("RouteView");return r(),d(p,{name:"zone-cp-detail-tabs-view",params:{zone:""}},{default:n(({can:z,route:k})=>[c(s,{src:`/zone-cps/${k.params.zone}`,onChange:v},{default:n(({data:y,error:V})=>[V!==void 0?(r(),d($,{key:0,error:V},null,8,["error"])):y===void 0?(r(),d(L,{key:1})):(r(),d(f,{key:2,breadcrumbs:[{to:{name:"zone-cp-list-view"},text:e(t)("zone-cps.routes.item.breadcrumbs")}]},j({title:n(()=>[g("h1",null,[c(P,{text:k.params.zone},{default:n(()=>[c(a,{title:e(t)("zone-cps.routes.item.title",{name:k.params.zone}),render:!0},null,8,["title"])]),_:2},1032,["text"])])]),default:n(()=>[_(),_(),c(Y,{class:"route-zone-detail-view-tabs",tabs:e(l)},null,8,["tabs"]),_(),c(m,null,{default:n(S=>[(r(),d(q(S.Component),{data:y,notifications:u.value},null,8,["data","notifications"]))]),_:2},1024)]),_:2},[z("create zones")?{name:"actions",fn:n(()=>[c(H,{"zone-overview":y},null,8,["zone-overview"])]),key:"0"}:void 0]),1032,["breadcrumbs"]))]),_:2},1032,["src"])]),_:1})}}});export{F as default};
