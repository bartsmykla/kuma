import{d as u,i as t,o as v,a as h,w as s,j as a,g as w,Y as b,k as n,P as f,J as V,t as x}from"./index-DKbsM-FP.js";const g=u({__name:"ServiceDetailTabsView",setup(R){return(T,A)=>{const r=t("RouteTitle"),m=t("XAction"),l=t("XTabs"),p=t("RouterView"),_=t("AppView"),d=t("RouteView");return v(),h(d,{name:"service-detail-tabs-view",params:{mesh:"",service:""}},{default:s(({route:e,t:i})=>[a(_,{breadcrumbs:[{to:{name:"mesh-detail-view",params:{mesh:e.params.mesh}},text:e.params.mesh},{to:{name:"service-list-view",params:{mesh:e.params.mesh}},text:i("services.routes.item.breadcrumbs")}]},{title:s(()=>[w("h1",null,[a(b,{text:e.params.service},{default:s(()=>[a(r,{title:i("services.routes.item.title",{name:e.params.service})},null,8,["title"])]),_:2},1032,["text"])])]),default:s(()=>{var c;return[n(),a(l,{selected:(c=e.child())==null?void 0:c.name},f({_:2},[V(e.children,({name:o})=>({name:`${o}-tab`,fn:s(()=>[a(m,{to:{name:o}},{default:s(()=>[n(x(i(`services.routes.item.navigation.${o}`)),1)]),_:2},1032,["to"])])}))]),1032,["selected"]),n(),a(p)]}),_:2},1032,["breadcrumbs"])]),_:1})}}});export{g as default};