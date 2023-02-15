import{k as w}from"./kumaApi-de9fdcba.js";import{u as x}from"./store-0511bcbf.js";import{_ as p}from"./CodeBlock.vue_vue_type_style_index_0_lang-be31e96b.js";import{L as y}from"./LoadingBox-74881ab8.js";import{O as N,a as A,b as C}from"./OnboardingPage-8b536a7f.js";import{d as T,r as _,c as P,l as D,a as E,w as i,o,e as s,f as e,u as I,h as a,F as L,g as t,b as O,p as S,j as R}from"./runtime-dom.esm-bundler-a6f4ece5.js";import{_ as B}from"./_plugin-vue_export-helper-c27b6911.js";import"./production-0f1ffdb6.js";import"./kongponents.es-8abed680.js";const h=c=>(S("data-v-c3f0ade0"),c=c(),R(),c),V=h(()=>t("p",{class:"mb-4 text-center"},`
        The demo application includes two services: a Redis backend to store a counter value, and a frontend web UI to show and increment the counter.
      `,-1)),G=h(()=>t("p",null,"To run execute the following command:",-1)),q={key:1},F={class:"status-box mt-4"},H={key:0,class:"status--is-connected","data-testid":"dpps-connected"},K={key:1,class:"status--is-disconnected","data-testid":"dpps-disconnected"},M={key:0,class:"status-loading-box mt-4"},U=T({__name:"AddNewServicesCode",setup(c){const f=x(),b=1e3,l="https://github.com/kumahq/kuma-counter-demo/",g="https://github.com/kumahq/kuma-counter-demo/blob/master/README.md",v="kubectl apply -f https://bit.ly/3Kh2Try",n=_(!1),r=_(null),k=P(()=>f.getters["config/getEnvironment"]==="kubernetes");u(),D(function(){m()});async function u(){try{const{total:d}=await w.getAllDataplanes();n.value=d>0}catch(d){console.error(d)}finally{n.value||(m(),r.value=window.setTimeout(()=>u(),b))}}function m(){r.value!==null&&window.clearTimeout(r.value)}return(d,$)=>(o(),E(N,null,{header:i(()=>[s(A,null,{title:i(()=>[e(`
          Add services
        `)]),_:1})]),content:i(()=>[V,e(),I(k)?(o(),a(L,{key:0},[G,e(),s(p,{id:"code-block-kubernetes-command",language:"bash",code:v})],64)):(o(),a("div",q,[t("p",{class:"mb-4 text-center"},[e(`
          Clone `),t("a",{href:l,target:"_blank"},"the GitHub repository"),e(` for the demo application:
        `)]),e(),s(p,{id:"code-block-clone-command",language:"bash",code:`git clone ${l}`},null,8,["code"]),e(),t("p",{class:"mt-4 text-center"},[e(`
          And follow the instructions in `),t("a",{href:g,target:"_blank"},"the README"),e(`.
        `)])])),e(),t("div",null,[t("p",F,[e(`
          DPPs status:

          `),n.value?(o(),a("span",H,"Connected")):(o(),a("span",K,"Disconnected"))]),e(),n.value?O("",!0):(o(),a("div",M,[s(y)]))])]),navigation:i(()=>[s(C,{"next-step":"onboarding-dataplanes-overview","previous-step":"onboarding-add-services","should-allow-next":n.value},null,8,["should-allow-next"])]),_:1}))}});const te=B(U,[["__scopeId","data-v-c3f0ade0"]]);export{te as default};
