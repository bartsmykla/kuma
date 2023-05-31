import{Q as O}from"./kongponents.es-a6f65032.js";import{L as A}from"./LoadingBox-06c57799.js";import{O as B,a as S,b as T}from"./OnboardingPage-2659dc4a.js";import{S as F}from"./StatusBadge-bae1de76.js";import{u as N}from"./index-8211ac1a.js";import{d as E,r as h,c as b,Q as I,R as L,o as s,b as y,w as t,g as n,y as Q,i as o,t as u,j as c,h as w,F as R,e as C}from"./index-ce1580f9.js";import{_ as H}from"./_plugin-vue_export-helper-c27b6911.js";import"./store-c50a9824.js";const M={key:0,class:"status-loading-box mb-4"},V={key:1},j={class:"mb-4"},z=E({__name:"DataplanesOverview",setup(K){const d=N(),k=[{label:"Mesh",key:"mesh"},{label:"Name",key:"name"},{label:"Status",key:"status"}],e=h({total:0,data:[]}),l=h(null),x=b(()=>e.value.data.length>0?"Success":"Waiting for DPPs"),p=b(()=>e.value.data.length>0?"The following data plane proxies (DPPs) are connected to the control plane:":null);I(function(){m()}),v();function m(){l.value!==null&&window.clearTimeout(l.value)}async function v(){let i=!1;const r=[];try{const{items:a}=await d.getAllDataplanes({size:10});if(Array.isArray(a))for(const D of a){const{name:f,mesh:_}=D,P=await d.getDataplaneOverviewFromMesh({mesh:_,name:f}),g=L(P.dataplaneInsight);g==="offline"&&(i=!0),r.push({status:g,name:f,mesh:_})}}catch(a){console.error(a)}e.value.data=r,e.value.total=e.value.data.length,i&&(m(),l.value=window.setTimeout(v,1e3))}return(i,r)=>(s(),y(T,null,{header:t(()=>[n(B,null,Q({title:t(()=>[o("p",null,u(x.value),1)]),_:2},[p.value!==null?{name:"description",fn:t(()=>[o("p",null,u(p.value),1)]),key:"0"}:void 0]),1024)]),content:t(()=>[e.value.data.length===0?(s(),c("div",M,[n(A)])):(s(),c("div",V,[o("p",j,[o("b",null,"Found "+u(e.value.data.length)+" DPPs:",1)]),w(),n(C(O),{class:"mb-4",fetcher:()=>e.value,headers:k,"disable-pagination":""},{status:t(({rowValue:a})=>[a?(s(),y(F,{key:0,status:a},null,8,["status"])):(s(),c(R,{key:1},[w(`
              —
            `)],64))]),_:1},8,["fetcher"])]))]),navigation:t(()=>[n(S,{"next-step":"onboarding-completed","previous-step":"onboarding-add-services-code","should-allow-next":e.value.data.length>0},null,8,["should-allow-next"])]),_:1}))}});const $=H(z,[["__scopeId","data-v-9ed5a755"]]);export{$ as default};
