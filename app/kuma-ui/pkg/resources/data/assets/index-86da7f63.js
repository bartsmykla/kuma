import{d as b,c as M,u as t,o as r,a as g,b as A,w as p,e as f,f as i,g as n,h as _,t as h,r as U,i as z,p as E,j as P,k as q,l as Q,n as Z,F as S,m as V,q as D,s as x,v as J,x as X,y as W,T as ee,z as te}from"./runtime-dom.esm-bundler-a6f4ece5.js";import{c as se,a as ne,u as K,b as Y}from"./vue-router-cf3250ac.js";import{s as O,u as $}from"./store-0511bcbf.js";import{C as B,c as oe,P as ae,g as ie,d as ce,a as w,T as N}from"./production-0f1ffdb6.js";import{q as re,m as C,T as le,L as I,O as G,n as ue,S as H,p as _e,_ as de,a as pe}from"./kongponents.es-8abed680.js";import{k as F}from"./kumaApi-de9fdcba.js";import{u as me,a as fe}from"./index-46eb6882.js";import{u as L,a as he}from"./index-28f79c9b.js";import{_ as k}from"./_plugin-vue_export-helper-c27b6911.js";import{d as ge}from"./datadogLogEvents-302eea7b.js";import{A as ve,a as ye}from"./AccordionItem-e30910c8.js";import"./DoughnutChart-ffc86670.js";(function(){const e=document.createElement("link").relList;if(e&&e.supports&&e.supports("modulepreload"))return;for(const c of document.querySelectorAll('link[rel="modulepreload"]'))l(c);new MutationObserver(c=>{for(const o of c)if(o.type==="childList")for(const d of o.addedNodes)d.tagName==="LINK"&&d.rel==="modulepreload"&&l(d)}).observe(document,{childList:!0,subtree:!0});function s(c){const o={};return c.integrity&&(o.integrity=c.integrity),c.referrerpolicy&&(o.referrerPolicy=c.referrerpolicy),c.crossorigin==="use-credentials"?o.credentials="include":c.crossorigin==="anonymous"?o.credentials="omit":o.credentials="same-origin",o}function l(c){if(c.ep)return;c.ep=!0;const o=s(c);fetch(c.href,o)}})();function be(a,e="/"){const s=se({history:ne(e),routes:a});return s.beforeEach(Me),s.beforeEach(Ae),s.beforeEach(ke),s}const Me=function(a,e,s){a.fullPath.startsWith("/#/")?s(a.fullPath.substring(2)):s()},Ae=function(a,e,s){a.params.mesh&&a.params.mesh!==O.state.selectedMesh&&O.dispatch("updateSelectedMesh",a.params.mesh),s()},ke=function(a,e,s){const l=O.state.onboarding.isCompleted,c=a.meta.onboardingProcess,o=O.getters.shouldSuggestOnboarding;l&&c&&!o?s({name:"home"}):!l&&!c&&o?s({name:B.get("onboardingStep")??"onboarding-welcome"}):s()},$e=b({__name:"AppBreadcrumbs",setup(a){const e=K(),s=Y(),l=M(()=>{const c=new Map;for(const o of e.matched){if(o.name==="home"||o.meta.parent==="home")continue;if(o.meta.parent!==void 0){const u=s.resolve({name:o.meta.parent});u.name&&c.set(u.name,{to:u,key:u.name,title:u.meta.title,text:u.meta.title})}if((o.name===e.name||o.redirect===e.name)&&o.meta.breadcrumbExclude!==!0&&e.name){let u=e.meta.title;e.meta.breadcrumbTitleParam&&e.params[e.meta.breadcrumbTitleParam]&&(u=e.params[e.meta.breadcrumbTitleParam]),c.set(e.name,{to:e,key:e.name,title:u,text:u})}}return Array.from(c.values())});return(c,o)=>t(l).length>0?(r(),g(t(re),{key:0,items:t(l)},null,8,["items"])):A("",!0)}}),Se=n("p",null,"Unable to reach the API",-1),we={key:0},Ne=b({__name:"AppErrorMessage",setup(a){return(e,s)=>(r(),g(t(le),{class:"global-api-status empty-state--wide-content empty-state--compact","cta-is-hidden":""},{title:p(()=>[f(t(C),{class:"mb-3",icon:"warning",color:"var(--black-75)","secondary-color":"var(--yellow-300)",size:"64"}),i(),Se]),message:p(()=>[n("p",null,[i(`
        Please double check to make sure it is up and running `),t(F).baseUrl?(r(),_("span",we,[i(", and it is reachable at "),n("code",null,h(t(F).baseUrl),1)])):A("",!0)])]),_:1}))}}),Ue={key:0,"data-testid":"notification-amount",class:"notification-icon__amount"},Ie=b({__name:"NotificationIcon",setup(a){const e=$(),s=M(()=>e.getters["notifications/amountOfActions"]);function l(){e.dispatch("notifications/openModal")}return(c,o)=>(r(),_("button",{class:"notification-icon cursor-pointer",type:"button",onClick:l},[f(t(C),{icon:"notificationBell",color:"var(--yellow-300)"}),i(),t(s)>0?(r(),_("span",Ue,h(t(s)),1)):A("",!0)]))}});const Le=k(Ie,[["__scopeId","data-v-cadae07a"]]),Re={class:"upgrade-check"},Ce={class:"alert-content"},Te=b({__name:"UpgradeCheck",setup(a){const e=L(),s=U(""),l=U(!1);c();async function c(){try{s.value=await F.getLatestVersion()}catch(o){l.value=!1,console.error(o)}finally{if(s.value!==""){const o=oe(s.value,e("KUMA_VERSION"));l.value=o===1}else{const d=new Date,u=new Date("2020-06-03 12:00:00"),v=new Date(u.getFullYear(),u.getMonth()+3,u.getDate());l.value=d.getTime()>=v.getTime()}}}return(o,d)=>(r(),_("div",Re,[l.value?(r(),g(t(G),{key:0,class:"upgrade-check-alert",appearance:"warning",size:"small"},{alertMessage:p(()=>[n("div",Ce,[n("div",null,h(t(e)("KUMA_PRODUCT_NAME"))+` update available
          `,1),i(),n("div",null,[f(t(I),{class:"warning-button",appearance:"primary",size:"small",to:t(e)("KUMA_INSTALL_URL")},{default:p(()=>[i(`
              Update
            `)]),_:1},8,["to"])])])]),_:1})):A("",!0)]))}});const xe=k(Te,[["__scopeId","data-v-03ce650c"]]),Oe=a=>(E("data-v-c348c723"),a=a(),P(),a),Ee={class:"app-header"},Pe={class:"horizontal-list"},Ke={class:"upgrade-check-wrapper"},De={key:0,class:"horizontal-list"},Be={class:"app-status app-status--mobile"},Fe={class:"app-status app-status--desktop"},ze=["href"],qe=["href"],Ve=Oe(()=>n("span",{class:"visually-hidden"},"Diagnostics",-1)),Ye=b({__name:"AppHeader",setup(a){const[e,s]=[me(),fe()],l=$(),c=L(),o=M(()=>l.getters["notifications/amountOfActions"]>0),d=M(()=>{const v=l.getters["config/getEnvironment"];return v?v.charAt(0).toUpperCase()+v.substring(1):"Universal"}),u=M(()=>l.getters["config/getMulticlusterStatus"]?"Multi-Zone":"Standalone");return(v,m)=>{const y=z("router-link");return r(),_("header",Ee,[n("div",Pe,[f(y,{to:{name:"home"}},{default:p(()=>[f(t(e))]),_:1}),i(),f(t(s),{class:"gh-star",href:"https://github.com/kumahq/kuma","aria-label":"Star kumahq/kuma on GitHub"},{default:p(()=>[i(`
        Star
      `)]),_:1}),i(),n("div",Ke,[f(xe)])]),i(),t(l).state.config.status==="OK"?(r(),_("div",De,[n("div",Be,[f(t(ue),{width:"280"},{content:p(()=>[n("p",null,[i(h(t(l).state.config.tagline)+" ",1),n("b",null,h(t(l).state.config.version),1),i(" on "),n("b",null,h(t(d)),1),i(" ("+h(t(u))+`)
            `,1)])]),default:p(()=>[f(t(I),{appearance:"outline"},{default:p(()=>[i(`
            Info
          `)]),_:1}),i()]),_:1})]),i(),n("p",Fe,[i(h(t(l).state.config.tagline)+" ",1),n("b",null,h(t(l).state.config.version),1),i(" on "),n("b",null,h(t(d)),1),i(" ("+h(t(u))+`)
      `,1)]),i(),t(o)?(r(),g(Le,{key:0})):A("",!0),i(),f(t(_e),{class:"help-menu",icon:"help","button-appearance":"outline","kpop-attributes":{placement:"bottomEnd"}},{items:p(()=>[f(t(H),null,{default:p(()=>[n("a",{href:`${t(c)("KUMA_DOCS_URL")}/?${t(c)("KUMA_UTM_QUERY_PARAMS")}`,target:"_blank",rel:"noopener noreferrer"},`
              Documentation
            `,8,ze)]),_:1}),i(),f(t(H),null,{default:p(()=>[n("a",{href:t(c)("KUMA_FEEDBACK_URL"),target:"_blank",rel:"noopener noreferrer"},`
              Feedback
            `,8,qe)]),_:1})]),_:1}),i(),f(t(I),{to:{name:"diagnostics"},icon:"gearFilled","button-appearance":"btn-link"},{icon:p(()=>[f(t(C),{icon:"gearFilled",class:"k-button-icon",size:"16",color:"currentColor","hide-title":""})]),default:p(()=>[i(),Ve]),_:1})])):A("",!0)])}}});const Ge=k(Ye,[["__scopeId","data-v-c348c723"]]),He=""+new URL("kuma-loader-v1-2aaed7d4.gif",import.meta.url).href,Qe=a=>(E("data-v-06e19708"),a=a(),P(),a),We={class:"full-screen"},je={class:"loading-container"},Ze=Qe(()=>n("img",{src:He},null,-1)),Je={class:"progress"},Xe=b({__name:"AppLoadingBar",setup(a){let e;const s=U(10);return q(function(){e=window.setInterval(()=>{s.value>=100&&(window.clearInterval(e),s.value=100),s.value=Math.min(s.value+Math.ceil(Math.random()*30),100)},150)}),Q(function(){window.clearInterval(e)}),(l,c)=>(r(),_("div",We,[n("div",je,[Ze,i(),n("div",Je,[n("div",{style:Z({width:`${s.value}%`}),class:"progress-bar",role:"progressbar","data-testid":"app-progress-bar"},null,4)])])]))}});const et=k(Xe,[["__scopeId","data-v-06e19708"]]),tt={key:0,class:"onboarding-check"},st={class:"alert-content"},nt=b({__name:"AppOnboardingNotification",setup(a){const e=U(!1);function s(){e.value=!0}return(l,c)=>e.value===!1?(r(),_("div",tt,[f(t(G),{appearance:"success",class:"dismissible","dismiss-type":"icon",onClosed:s},{alertMessage:p(()=>[n("div",st,[n("div",null,[n("strong",null,"Welcome to "+h(t(ae))+"!",1),i(` We've detected that you don't have any data plane proxies running yet. We've created an onboarding process to help you!
          `)]),i(),n("div",null,[f(t(I),{appearance:"primary",size:"small",class:"action-button",to:{name:"onboarding-welcome"}},{default:p(()=>[i(`
              Get started
            `)]),_:1})])])]),_:1})])):A("",!0)}});const ot=k(nt,[["__scopeId","data-v-c21dc5a7"]]);async function at(a,e,s=()=>!1){do{if(await a(),await s())break;const l=typeof e=="number"?e:e();await new Promise(c=>setTimeout(c,Math.max(0,l)))}while(!await s())}const it=a=>(E("data-v-76b8351f"),a=a(),P(),a),ct={class:"mesh-selector-container"},rt={for:"mesh-selector"},lt=it(()=>n("span",{class:"visually-hidden"},`
        Filter by mesh:
      `,-1)),ut=["value","selected"],_t=b({__name:"AppMeshSelector",props:{meshes:{type:Array,required:!0}},setup(a){const e=a,s=K(),l=Y(),c=$(),o=M(()=>c.state.selectedMesh===null?e.meshes[0].name:c.state.selectedMesh);function d(u){const m=u.target.value;c.dispatch("updateSelectedMesh",m);const y="mesh"in s.params?s.name:"mesh-detail-view";l.push({name:y,params:{mesh:m}})}return(u,v)=>(r(),_("div",ct,[n("label",rt,[lt,i(),n("select",{id:"mesh-selector",class:"mesh-selector",name:"mesh-selector","data-testid":"mesh-selector",onChange:d},[(r(!0),_(S,null,V(e.meshes,m=>(r(),_("option",{key:m.name,value:m.name,selected:m.name===t(o)},h(m.name),9,ut))),128))],32)])]))}});const dt=k(_t,[["__scopeId","data-v-76b8351f"]]),pt=["data-testid"],mt={key:1,class:"nav-category"},ft=b({__name:"AppNavItem",props:{name:{type:String,required:!0},routeName:{type:String,required:!1,default:""},usesMeshParam:{type:Boolean,required:!1,default:!1},categoryTier:{type:String,required:!1,default:null},insightsFieldAccessor:{type:String,required:!1,default:""},shouldOffsetFromFollowingItems:{type:Boolean,required:!1,default:!1}},setup(a){const e=a,s=K(),l=Y(),c=$(),o=M(()=>{if(e.insightsFieldAccessor){const m=ie(c.state.sidebar.insights,e.insightsFieldAccessor,0);return m>99?"99+":String(m)}else return""}),d=M(()=>{if(e.routeName==="")return null;const m={name:e.routeName};return e.usesMeshParam&&(m.params={mesh:c.state.selectedMesh}),m}),u=M(()=>{if(d.value===null)return!1;if(e.routeName===s.name||s.path.split("/")[2]===d.value.name)return!0;if(s.meta.parent)try{if(l.resolve({name:s.meta.parent}).name===e.routeName)return!0}catch(y){if(y instanceof Error&&y.message.includes("No match for"))console.warn(y);else throw y}return e.routeName&&s.matched.some(y=>e.routeName===y.name||e.routeName===y.redirect)});function v(){ce.logger.info(ge.SIDEBAR_ITEM_CLICKED,{data:d.value})}return(m,y)=>{const R=z("router-link");return r(),_("div",{class:D(["nav-item",{"nav-item--is-category":t(d)===null,"nav-item--has-bottom-offset":e.shouldOffsetFromFollowingItems,[`nav-item--is-${e.categoryTier}-category`]:e.categoryTier!==null}]),"data-testid":e.routeName},[t(d)!==null?(r(),g(R,{key:0,class:D(["nav-link",{"nav-link--is-active":t(u)}]),to:t(d),onClick:v},{default:p(()=>[i(h(a.name)+" ",1),t(o)?(r(),_("span",{key:0,class:D(["amount",{"amount--empty":t(o)==="0"}])},h(t(o)),3)):A("",!0)]),_:1},8,["class","to"])):(r(),_("div",mt,h(a.name),1))],10,pt)}}});const ht=k(ft,[["__scopeId","data-v-938e565b"]]),gt={class:"app-sidebar-wrapper"},vt={class:"app-sidebar"},yt=b({__name:"AppSidebar",setup(a){const s=$(),l=he(),c=M(()=>l(s.getters["config/getMulticlusterStatus"],s.state.meshes.items.length>0)),o=M(()=>s.state.meshes.items);x(()=>s.state.selectedMesh,()=>{s.dispatch("sidebar/getMeshInsights")});let d=!1;q(function(){window.addEventListener("blur",u),window.addEventListener("focus",v)}),Q(function(){window.removeEventListener("blur",u),window.removeEventListener("focus",v)}),v();function u(){d=!0}function v(){d=!1,at(m,10*1e3,()=>d)}function m(){return s.dispatch("sidebar/getInsights")}return(y,R)=>(r(),_("div",gt,[n("aside",vt,[(r(!0),_(S,null,V(t(c),(T,j)=>(r(),_(S,{key:j},[T.isMeshSelector?(r(),_(S,{key:0},[t(o).length>0?(r(),g(dt,{key:0,meshes:t(o)},null,8,["meshes"])):A("",!0)],64)):(r(),g(ht,J(X({key:1},T)),null,16))],64))),128))])]))}});const bt=k(yt,[["__scopeId","data-v-ddb44585"]]),Mt={class:"py-4"},At=n("p",{class:"mb-4"},`
      A traffic log policy lets you collect access logs for every data plane proxy in your service mesh.
    `,-1),kt={class:"list-disc pl-4"},$t=["href"],St=b({__name:"LoggingNotification",setup(a){const e=L();return(s,l)=>(r(),_("div",Mt,[At,i(),n("ul",kt,[n("li",null,[n("a",{href:`${t(e)("KUMA_DOCS_URL")}/policies/traffic-log/?${t(e)("KUMA_UTM_QUERY_PARAMS")}`,target:"_blank"},`
          Traffic Log policy documentation
        `,8,$t)])])]))}}),wt={class:"py-4"},Nt=n("p",{class:"mb-4"},`
      A traffic metrics policy lets you collect key data for observability of your service mesh.
    `,-1),Ut={class:"list-disc pl-4"},It=["href"],Lt=b({__name:"MetricsNotification",setup(a){const e=L();return(s,l)=>(r(),_("div",wt,[Nt,i(),n("ul",Ut,[n("li",null,[n("a",{href:`${t(e)("KUMA_DOCS_URL")}/policies/traffic-metrics/?${t(e)("KUMA_UTM_QUERY_PARAMS")}`,target:"_blank"},`
          Traffic Metrics policy documentation
        `,8,It)])])]))}}),Rt={class:"py-4"},Ct=n("p",{class:"mb-4"},`
      Mutual TLS (mTLS) for communication between all the components
      of your service mesh (services, control plane, data plane proxies), proxy authentication,
      and access control rules in Traffic Permissions policies all contribute to securing your mesh.
    `,-1),Tt={class:"list-disc pl-4"},xt=["href"],Ot=["href"],Et=["href"],Pt=b({__name:"MtlsNotification",setup(a){const e=L();return(s,l)=>(r(),_("div",Rt,[Ct,i(),n("ul",Tt,[n("li",null,[n("a",{href:`${t(e)("KUMA_DOCS_URL")}/security/certificates/?${t(e)("KUMA_UTM_QUERY_PARAMS")}`,target:"_blank"},`
          Secure access across services
        `,8,xt)]),i(),n("li",null,[n("a",{href:`${t(e)("KUMA_DOCS_URL")}/policies/mutual-tls/?${t(e)("KUMA_UTM_QUERY_PARAMS")}`,target:"_blank"},`
          Mutual TLS
        `,8,Ot)]),i(),n("li",null,[n("a",{href:`${t(e)("KUMA_DOCS_URL")}/policies/traffic-permissions/?${t(e)("KUMA_UTM_QUERY_PARAMS")}`,target:"_blank"},`
          Traffic Permissions policy documentation
        `,8,Et)])])]))}}),Kt={class:"py-4"},Dt=n("p",{class:"mb-4"},`
      A traffic trace policy lets you enable tracing logs and a third-party tracing solution to send them to.
    `,-1),Bt={class:"list-disc pl-4"},Ft=["href"],zt=b({__name:"TracingNotification",setup(a){const e=L();return(s,l)=>(r(),_("div",Kt,[Dt,i(),n("ul",Bt,[n("li",null,[n("a",{href:`${t(e)("KUMA_DOCS_URL")}/policies/traffic-trace/?${t(e)("KUMA_UTM_QUERY_PARAMS")}`,target:"_blank"},`
          Traffic Trace policy documentation
        `,8,Ft)])])]))}}),qt={class:"flex items-center"},Vt=b({__name:"SingleMeshNotifications",setup(a){const e=$(),s={LoggingNotification:St,MetricsNotification:Lt,MtlsNotification:Pt,TracingNotification:zt};return(l,c)=>(r(),g(ye,{"multiple-open":""},{default:p(()=>[(r(!0),_(S,null,V(t(e).getters["notifications/singleMeshNotificationItems"],o=>(r(),g(ve,{key:o.name},{"accordion-header":p(()=>[n("div",qt,[o.isCompleted?(r(),g(t(C),{key:0,color:"var(--green-500)",icon:"check",size:"20",class:"mr-4"})):(r(),g(t(C),{key:1,icon:"warning",color:"var(--black-75)","secondary-color":"var(--yellow-300)",size:"20",class:"mr-4"})),i(),n("strong",null,h(o.name),1)])]),"accordion-content":p(()=>[o.component?(r(),g(W(s[o.component]),{key:0})):(r(),g(t(de),{key:1},{body:p(()=>[i(h(o.content),1)]),_:2},1024))]),_:2},1024))),128))]),_:1}))}}),Yt=a=>(E("data-v-baf26e82"),a=a(),P(),a),Gt={class:"mr-4"},Ht=Yt(()=>n("span",{class:"mr-2"},[n("strong",null,"Pro tip:"),i(`

            You might want to adjust your mesh configuration
          `)],-1)),Qt={key:0},Wt={class:"text-xl tracking-wide"},jt={key:1},Zt={class:"text-xl tracking-wide"},Jt=b({__name:"NotificationManager",setup(a){const e=$(),s=U(!0),l=M(()=>e.state.selectedMesh?e.getters["notifications/meshNotificationItemMapWithAction"][e.state.selectedMesh]:!1);q(function(){const u=B.get("hideCheckMeshAlert");s.value=u!=="yes"});function c(){s.value=!1,B.set("hideCheckMeshAlert","yes")}function o(){e.dispatch("notifications/openModal")}function d(){e.dispatch("notifications/closeModal")}return(u,v)=>(r(),_("div",null,[s.value?(r(),g(t(G),{key:0,class:"mb-4",appearance:"info","dismiss-type":"icon","data-testid":"notification-info",onClosed:c},{alertMessage:p(()=>[n("div",Gt,[Ht,i(),f(t(I),{appearance:"outline","data-testid":"open-modal-button",onClick:o},{default:p(()=>[i(`
            Check your mesh!
          `)]),_:1})])]),_:1})):A("",!0),i(),f(t(pe),{class:"modal","is-visible":t(e).state.notifications.isOpen,title:"Notifications","text-align":"left","data-testid":"notification-modal"},{"header-content":p(()=>[n("div",null,[n("div",null,[t(l)?(r(),_("span",Qt,[i(`
              Some of these features are not enabled for `),n("span",Wt,'"'+h(t(e).state.selectedMesh)+'"',1),i(` mesh. Consider implementing them.
            `)])):(r(),_("span",jt,[i(`
              Looks like `),n("span",Zt,'"'+h(t(e).state.selectedMesh)+'"',1),i(` isn't missing any features. Well done!
            `)]))])])]),"body-content":p(()=>[f(Vt)]),"footer-content":p(()=>[f(t(I),{appearance:"outline","data-testid":"close-modal-button",onClick:d},{default:p(()=>[i(`
          Close
        `)]),_:1})]),_:1},8,["is-visible"])]))}});const Xt=k(Jt,[["__scopeId","data-v-baf26e82"]]),es={key:0},ts={key:1,class:"app-content-container"},ss={class:"app-main-content"},ns=b({__name:"App",setup(a){const e=$(),s=K(),l=U(e.state.globalLoading),c=M(()=>s.path),o=M(()=>e.state.config.status!=="OK"),d=M(()=>e.getters.shouldSuggestOnboarding),u=M(()=>e.getters["notifications/amountOfActions"]>0);x(()=>e.state.globalLoading,function(m){l.value=m}),x(()=>s.meta.title,function(m){v(m)}),x(()=>e.state.pageTitle,function(m){v(m)});function v(m){const y="Kuma Manager";document.title=m?`${m} | ${y}`:y}return(m,y)=>{const R=z("router-view");return l.value?(r(),g(et,{key:0})):(r(),_(S,{key:1},[f(Ge),i(),t(s).meta.onboardingProcess?(r(),_("div",es,[f(R)])):(r(),_("div",ts,[f(bt),i(),n("main",ss,[t(o)?(r(),g(Ne,{key:0})):A("",!0),i(),t(u)?(r(),g(Xt,{key:1})):A("",!0),i(),t(d)?(r(),g(ot,{key:2})):A("",!0),i(),f($e),i(),(r(),g(R,{key:t(c)},{default:p(({Component:T})=>[f(ee,{mode:"out-in",name:"fade"},{default:p(()=>[(r(),_("div",{key:t(s).name,class:"transition-root"},[(r(),g(W(T)))]))]),_:2},1024)]),_:1}))])]))],64))}}});const os=k(ns,[["__scopeId","data-v-50928ee3"]]);async function as(a,e,s){const l=w(N.store),c=w(N.api);document.title=`${a("KUMA_PRODUCT_NAME")} Manager`,c.setBaseUrl(a("KUMA_API_URL")),(async()=>{const u=await c.getConfig();s.setup(u)})();const o=te(os);o.use(l,w(N.storeKey)),await Promise.all([l.dispatch("bootstrap"),l.dispatch("fetchPolicyTypes")]);const d=await be(e,a("KUMA_BASE_PATH"));o.use(d),o.mount("#app")}as(w(N.env),w(N.routes),w(N.logger));
