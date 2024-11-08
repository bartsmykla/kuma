import{d as z,e as y,o as s,m as p,w as t,a as c,k as _,X as l,b as e,t as a,c as d,H as u,J as m,p as h,l as V,af as I,r as A}from"./index-BWlxH9e6.js";const C={class:"stack-with-borders"},N={key:1,class:"mt-8 stack-with-borders"},B=z({__name:"SubscriptionSummaryOverviewView",props:{data:{},routeName:{}},setup(b){const n=b;return(k,S)=>{const f=y("XAlert"),w=y("AppView"),v=y("RouteView");return s(),p(v,{name:n.routeName},{default:t(({t:r})=>[c(w,null,{default:t(()=>[_("div",C,[c(l,{layout:"horizontal"},{title:t(()=>[e(a(r("http.api.property.version")),1)]),body:t(()=>{var o,i;return[(s(!0),d(u,null,m([(i=(o=n.data.version)==null?void 0:o.kumaCp)==null?void 0:i.version],g=>(s(),d(u,null,[e(a(g??"-"),1)],64))),256))]}),_:2},1024),e(),c(l,{layout:"horizontal"},{title:t(()=>[e(a(r("http.api.property.connectTime")),1)]),body:t(()=>[e(a(r("common.formats.datetime",{value:Date.parse(n.data.connectTime??"")})),1)]),_:2},1024),e(),n.data.disconnectTime?(s(),p(l,{key:0,layout:"horizontal"},{title:t(()=>[e(a(r("http.api.property.disconnectTime")),1)]),body:t(()=>[e(a(r("common.formats.datetime",{value:Date.parse(n.data.disconnectTime)})),1)]),_:2},1024)):h("",!0),e(),c(l,{layout:"horizontal"},{title:t(()=>[e(a(r("subscriptions.routes.item.headers.responses")),1)]),body:t(()=>{var o;return[(s(!0),d(u,null,m([((o=n.data.status)==null?void 0:o.total)??{}],i=>(s(),d(u,null,[e(a(i.responsesSent)+"/"+a(i.responsesAcknowledged),1)],64))),256))]}),_:2},1024),e(),(s(),d(u,null,m(["zoneInstanceId","globalInstanceId","controlPlaneInstanceId"],o=>(s(),d(u,{key:typeof o},[n.data[o]?(s(),p(l,{key:0,layout:"horizontal"},{title:t(()=>[e(a(r(`http.api.property.${o}`)),1)]),body:t(()=>[e(a(n.data[o]),1)]),_:2},1024)):h("",!0)],64))),64)),e(),c(l,{layout:"horizontal"},{title:t(()=>[e(a(r("http.api.property.id")),1)]),body:t(()=>[e(a(n.data.id),1)]),_:2},1024)]),e(),Object.keys(n.data.status.acknowledgements).length===0?(s(),p(f,{key:0,appearance:"info"},{icon:t(()=>[c(V(I))]),default:t(()=>[e(" "+a(r("common.detail.subscriptions.no_stats",{id:n.data.id})),1)]),_:2},1024)):(s(),d("div",N,[_("div",null,[A(k.$slots,"default")]),e(),c(l,{class:"mt-4",layout:"horizontal"},{title:t(()=>[_("strong",null,a(r("subscriptions.routes.item.headers.type")),1)]),body:t(()=>[e(a(r("subscriptions.routes.item.headers.stat")),1)]),_:2},1024),e(),(s(!0),d(u,null,m(Object.entries(n.data.status.acknowledgements??{}),([o,i])=>(s(),p(l,{key:o,layout:"horizontal"},{title:t(()=>[e(a(r(`http.api.property.${o}`)),1)]),body:t(()=>[e(a(i.responsesSent)+"/"+a(i.responsesAcknowledged),1)]),_:2},1024))),128))]))]),_:2},1024)]),_:3},8,["name"])}}});export{B as default};
