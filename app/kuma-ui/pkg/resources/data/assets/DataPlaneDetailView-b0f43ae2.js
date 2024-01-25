import{a as J,K as A}from"./index-fce48c05.js";import{_ as q,a as y,o,b as f,w as e,m as n,r as S,f as t,d as U,c as d,e as r,k as j,t as l,l as a,F as p,C as b,p as $,n as Q,y as tt,H as et,X as W,Y as w,Z as at,$ as nt,a0 as st,O as ot,a1 as lt,B as rt,s as it,v as dt}from"./index-3ddd0e9e.js";import{_ as ct}from"./LoadingBlock.vue_vue_type_script_setup_true_lang-4407ccd5.js";import{S as ut}from"./StatusBadge-9883c335.js";import{S as _t}from"./SummaryView-53a1c27b.js";import{T as X}from"./TagList-8b3f124d.js";import{T as pt}from"./TextWithCopyButton-4870eafb.js";import{_ as mt}from"./SubscriptionList.vue_vue_type_script_setup_true_lang-0e9e752a.js";import"./CopyButton-f0ea0e69.js";import"./AccordionList-0e0b7f3f.js";const ft={},vt={class:"card"},yt={class:"title"},ht={class:"body"};function gt(h,s){const c=y("KCard");return o(),f(c,{class:"data-card"},{default:e(()=>[n("dl",null,[n("div",vt,[n("dt",yt,[S(h.$slots,"title",{},void 0,!0)]),t(),n("dd",ht,[S(h.$slots,"default",{},void 0,!0)])])])]),_:3})}const Y=q(ft,[["render",gt],["__scopeId","data-v-6e083223"]]),bt={class:"service-traffic"},kt={class:"actions"},xt=U({__name:"DataPlaneTraffic",setup(h){return(s,c)=>(o(),d("div",bt,[n("div",kt,[S(s.$slots,"actions",{},void 0,!0)]),t(),r(Y,{class:"header"},{title:e(()=>[S(s.$slots,"title",{},void 0,!0)]),_:3}),t(),S(s.$slots,"default",{},void 0,!0)]))}});const Z=q(xt,[["__scopeId","data-v-5bd1dbf9"]]),wt={class:"title"},$t={key:0},Ct=U({__name:"ServiceTrafficCard",props:{protocol:{},traffic:{}},setup(h){const{t:s}=j(),c=h,K=_=>{const T=_.target;if(T.nodeName.toLowerCase()!=="a"){const I=T.closest(".service-traffic-card");if(I){const D=I.querySelector("a");D!==null&&D.click()}}};return(_,T)=>{const I=y("KBadge"),D=y("KSkeletonBox");return o(),f(Y,{class:"service-traffic-card",onClick:K},{title:e(()=>[r(I,{appearance:c.protocol==="passthrough"?"success":"info"},{default:e(()=>[t(l(a(s)(`data-planes.components.service_traffic_card.protocol.${c.protocol}`,{},{defaultMessage:a(s)(`http.api.value.${c.protocol}`)})),1)]),_:1},8,["appearance"]),t(),n("div",wt,[S(_.$slots,"default",{},void 0,!0)])]),default:e(()=>{var z,C,V,B,E,N,L,M,O;return[t(),c.traffic?(o(),d(p,{key:0},[c.traffic.name.length>0?(o(),d("dl",$t,[c.protocol==="passthrough"?(o(!0),d(p,{key:0},b([["http","tcp"].reduce((m,R)=>{var g;const v="downstream";return Object.entries(((g=c.traffic)==null?void 0:g[R])||{}).reduce((x,[P,i])=>[`${v}_cx_tx_bytes_total`,`${v}_cx_rx_bytes_total`].includes(P)?{...x,[P]:i+(x[P]??0)}:x,m)},{})],(m,R)=>(o(),d(p,{key:R},[n("div",null,[n("dt",null,l(a(s)("data-planes.components.service_traffic_card.tx")),1),t(),n("dd",null,l(a(s)("common.formats.bytes",{value:m.downstream_cx_rx_bytes_total??0})),1)]),t(),n("div",null,[n("dt",null,l(a(s)("data-planes.components.service_traffic_card.rx")),1),t(),n("dd",null,l(a(s)("common.formats.bytes",{value:m.downstream_cx_tx_bytes_total??0})),1)])],64))),128)):c.protocol==="grpc"?(o(),d(p,{key:1},[n("div",null,[n("dt",null,l(a(s)("data-planes.components.service_traffic_card.grpc_success")),1),t(),n("dd",null,l(a(s)("common.formats.integer",{value:((z=c.traffic.grpc)==null?void 0:z.success)??0})),1)]),t(),n("div",null,[n("dt",null,l(a(s)("data-planes.components.service_traffic_card.grpc_failure")),1),t(),n("dd",null,l(a(s)("common.formats.integer",{value:((C=c.traffic.grpc)==null?void 0:C.failure)??0})),1)])],64)):c.protocol==="http"?(o(),d(p,{key:2},[(o(!0),d(p,null,b([((V=c.traffic.http)==null?void 0:V.downstream_rq_1xx)??0].filter(m=>m!==0),m=>(o(),d("div",{key:m},[n("dt",null,l(a(s)("data-planes.components.service_traffic_card.1xx")),1),t(),n("dd",null,l(a(s)("common.formats.integer",{value:m})),1)]))),128)),t(),n("div",null,[n("dt",null,l(a(s)("data-planes.components.service_traffic_card.2xx")),1),t(),n("dd",null,l(a(s)("common.formats.integer",{value:((B=c.traffic.http)==null?void 0:B.downstream_rq_2xx)??0})),1)]),t(),(o(!0),d(p,null,b([((E=c.traffic.http)==null?void 0:E.downstream_rq_3xx)??0].filter(m=>m!==0),m=>(o(),d("div",{key:m},[n("dt",null,l(a(s)("data-planes.components.service_traffic_card.3xx")),1),t(),n("dd",null,l(a(s)("common.formats.integer",{value:m})),1)]))),128)),t(),n("div",null,[n("dt",null,l(a(s)("data-planes.components.service_traffic_card.4xx")),1),t(),n("dd",null,l(a(s)("common.formats.integer",{value:((N=c.traffic.http)==null?void 0:N.downstream_rq_4xx)??0})),1)]),t(),n("div",null,[n("dt",null,l(a(s)("data-planes.components.service_traffic_card.5xx")),1),t(),n("dd",null,l(a(s)("common.formats.integer",{value:((L=c.traffic.http)==null?void 0:L.downstream_rq_5xx)??0})),1)])],64)):(o(),d(p,{key:3},[n("div",null,[n("dt",null,l(a(s)("data-planes.components.service_traffic_card.tx")),1),t(),n("dd",null,l(a(s)("common.formats.bytes",{value:((M=c.traffic.tcp)==null?void 0:M.downstream_cx_rx_bytes_total)??0})),1)]),t(),n("div",null,[n("dt",null,l(a(s)("data-planes.components.service_traffic_card.rx")),1),t(),n("dd",null,l(a(s)("common.formats.bytes",{value:((O=c.traffic.tcp)==null?void 0:O.downstream_cx_tx_bytes_total)??0})),1)])],64))])):$("",!0)],64)):(o(),f(D,{key:1,width:"10"}))]}),_:3})}}});const H=q(Ct,[["__scopeId","data-v-d859406b"]]),St={class:"body"},Tt=U({__name:"ServiceTrafficGroup",props:{type:{}},setup(h){const s=h;return(c,K)=>{const _=y("KCard");return o(),f(_,{class:Q(["service-traffic-group",`type-${s.type}`])},{default:e(()=>[n("div",St,[S(c.$slots,"default",{},void 0,!0)])]),_:3},8,["class"])}}});const G=q(Tt,[["__scopeId","data-v-baf4abf7"]]),Kt=h=>(it("data-v-5a7bf611"),h=h(),dt(),h),It={"data-testid":"dataplane-warnings"},Dt=["data-testid","innerHTML"],Vt={key:0,"data-testid":"warning-stats-loading"},Bt={class:"stack","data-testid":"dataplane-details"},Rt={class:"columns"},Pt={class:"status-with-reason"},qt={class:"columns"},zt=Kt(()=>n("span",null,"Outbounds",-1)),Et={"data-testid":"dataplane-mtls"},Nt={class:"columns"},Lt=["innerHTML"],Mt={key:1,"data-testid":"dataplane-subscriptions"},Ot=U({__name:"DataPlaneDetailView",props:{data:{}},setup(h){const{t:s,formatIsoDate:c}=j(),K=tt(),_=h,T=et(()=>_.data.warnings.concat(..._.data.isCertExpired?[{kind:"CERT_EXPIRED"}]:[]));return(I,D)=>{const z=y("KTooltip"),C=y("KCard"),V=y("RouterLink"),B=y("DataCollection"),E=y("KInputSwitch"),N=y("KButton"),L=y("RouterView"),M=y("KAlert"),O=y("AppView"),m=y("DataSource"),R=y("RouteView");return o(),f(R,{params:{mesh:"",dataPlane:"",inactive:!1},name:"data-plane-detail-view"},{default:e(({route:v})=>[r(m,{src:_.data.dataplaneType==="standard"?`/meshes/${v.params.mesh}/dataplanes/${v.params.dataPlane}/stats`:""},{default:e(({data:g,error:x,refresh:P})=>[r(O,null,W({default:e(()=>[t(),n("div",Bt,[r(C,null,{default:e(()=>[n("div",Rt,[r(w,null,{title:e(()=>[t(l(a(s)("http.api.property.status")),1)]),body:e(()=>[n("div",Pt,[r(ut,{status:_.data.status},null,8,["status"]),t(),(o(!0),d(p,null,b([_.data.dataplane.networking.inbounds.filter(i=>!i.health.ready)],i=>(o(),d(p,{key:i},[i.length>0?(o(),f(z,{key:0,class:"reason-tooltip","position-fixed":""},{content:e(()=>[n("ul",null,[(o(!0),d(p,null,b(i,u=>(o(),d("li",{key:`${u.service}:${u.port}`},l(a(s)("data-planes.routes.item.unhealthy_inbound",{service:u.service,port:u.port})),1))),128))])]),default:e(()=>[r(a(at),{color:a(J),size:a(A),"hide-title":""},null,8,["color","size"]),t()]),_:2},1024)):$("",!0)],64))),128))])]),_:1}),t(),r(w,null,{title:e(()=>[t(l(a(s)("data-planes.routes.item.last_updated")),1)]),body:e(()=>[t(l(a(c)(_.data.modificationTime)),1)]),_:1}),t(),_.data.dataplane.networking.gateway?(o(),d(p,{key:0},[r(w,null,{title:e(()=>[t(l(a(s)("http.api.property.tags")),1)]),body:e(()=>[r(X,{tags:_.data.dataplane.networking.gateway.tags},null,8,["tags"])]),_:1}),t(),r(w,null,{title:e(()=>[t(l(a(s)("http.api.property.address")),1)]),body:e(()=>[r(pt,{text:`${_.data.dataplane.networking.address}`},null,8,["text"])]),_:1})],64)):$("",!0)])]),_:1}),t(),_.data.dataplaneType==="standard"?(o(),f(C,{key:0,class:"traffic","data-testid":"dataplane-traffic"},{default:e(()=>[n("div",qt,[r(Z,null,{title:e(()=>[r(a(nt),{display:"inline-block",decorative:"",size:a(A)},null,8,["size"]),t(`
                  Inbounds
                `)]),default:e(()=>[t(),r(G,{type:"inbound"},{default:e(()=>[r(B,{items:_.data.dataplane.networking.inbounds},{default:e(({items:i})=>[(o(!0),d(p,null,b(i,u=>(o(),d(p,{key:`${u.port}`},[(o(!0),d(p,null,b([(g||{inbounds:[]}).inbounds.find(k=>`${k.port}`==`${u.port}`)],k=>(o(),f(H,{key:k,protocol:u.protocol,traffic:typeof x>"u"?k:{name:"",protocol:u.protocol,port:`${u.port}`}},{default:e(()=>[r(V,{to:{name:(F=>F.includes("bound")?F.replace("-outbound-","-inbound-"):"data-plane-inbound-summary-overview-view")(String(a(K).name)),params:{service:u.port},query:{inactive:v.params.inactive?null:void 0}}},{default:e(()=>[t(`
                            :`+l(u.port),1)]),_:2},1032,["to"]),t(),r(X,{tags:[{label:"kuma.io/service",value:u.tags["kuma.io/service"]}]},null,8,["tags"])]),_:2},1032,["protocol","traffic"]))),128))],64))),128))]),_:2},1032,["items"])]),_:2},1024)]),_:2},1024),t(),r(Z,null,W({title:e(()=>[r(a(st),{display:"inline-block",decorative:"",size:a(A)},null,8,["size"]),t(),zt]),default:e(()=>[t(),t(),typeof x>"u"?(o(),d(p,{key:0},[typeof g>"u"?(o(),f(ct,{key:0})):(o(),d(p,{key:1},[r(G,{type:"passthrough"},{default:e(()=>[r(H,{protocol:"passthrough",traffic:g.passthrough},{default:e(()=>[t(`
                        Non mesh traffic
                      `)]),_:2},1032,["traffic"])]),_:2},1024),t(),r(B,{predicate:v.params.inactive?void 0:i=>{var u,k;return((i.protocol==="tcp"?(u=i.tcp)==null?void 0:u.downstream_cx_rx_bytes_total:(k=i.http)==null?void 0:k.downstream_rq_total)??0)>0},items:g.outbounds},{default:e(({items:i})=>[i.length>0?(o(),f(G,{key:0,type:"outbound","data-testid":"dataplane-outbounds"},{default:e(()=>[(o(!0),d(p,null,b(i,u=>(o(),f(H,{key:`${u.name}`,protocol:u.protocol,traffic:u},{default:e(()=>[r(V,{to:{name:(k=>k.includes("bound")?k.replace("-inbound-","-outbound-"):"data-plane-outbound-summary-overview-view")(String(a(K).name)),params:{service:u.name},query:{inactive:v.params.inactive?null:void 0}}},{default:e(()=>[t(l(u.name),1)]),_:2},1032,["to"])]),_:2},1032,["protocol","traffic"]))),128))]),_:2},1024)):$("",!0)]),_:2},1032,["predicate","items"])],64))],64)):(o(),f(ot,{key:1}))]),_:2},[g?{name:"actions",fn:e(()=>[r(E,{modelValue:v.params.inactive,"onUpdate:modelValue":i=>v.params.inactive=i,"data-testid":"dataplane-outbounds-inactive-toggle"},{label:e(()=>[t(`
                      Show inactive
                    `)]),_:2},1032,["modelValue","onUpdate:modelValue"]),t(),r(N,{appearance:"primary",onClick:P},{default:e(()=>[r(a(lt),{size:a(A)},null,8,["size"]),t(`

                    Refresh
                  `)]),_:2},1032,["onClick"])]),key:"0"}:void 0]),1024)])]),_:2},1024)):$("",!0),t(),r(L,null,{default:e(i=>[i.route.name!==v.name?(o(),f(_t,{key:0,width:"670px",onClose:function(u){v.replace({name:"data-plane-detail-view",params:{mesh:v.params.mesh,dataPlane:v.params.dataPlane},query:{inactive:v.params.inactive?null:void 0}})}},{default:e(()=>[(o(),f(rt(i.Component),{data:String(i.route.name).includes("-inbound-")?_.data.dataplane.networking.inbounds||[]:(g==null?void 0:g.outbounds)||[]},null,8,["data"]))]),_:2},1032,["onClose"])):$("",!0)]),_:2},1024),t(),n("div",Et,[n("h2",null,l(a(s)("data-planes.routes.item.mtls.title")),1),t(),_.data.dataplaneInsight.mTLS?(o(!0),d(p,{key:0},b([_.data.dataplaneInsight.mTLS],i=>(o(),f(C,{key:i,class:"mt-4"},{default:e(()=>[n("div",Nt,[r(w,null,{title:e(()=>[t(l(a(s)("data-planes.routes.item.mtls.expiration_time.title")),1)]),body:e(()=>[t(l(a(c)(i.certificateExpirationTime)),1)]),_:2},1024),t(),r(w,null,{title:e(()=>[t(l(a(s)("data-planes.routes.item.mtls.generation_time.title")),1)]),body:e(()=>[t(l(a(c)(i.lastCertificateRegeneration)),1)]),_:2},1024),t(),r(w,null,{title:e(()=>[t(l(a(s)("data-planes.routes.item.mtls.regenerations.title")),1)]),body:e(()=>[t(l(a(s)("common.formats.integer",{value:i.certificateRegenerations})),1)]),_:2},1024),t(),r(w,null,{title:e(()=>[t(l(a(s)("data-planes.routes.item.mtls.issued_backend.title")),1)]),body:e(()=>[t(l(i.issuedBackend),1)]),_:2},1024),t(),r(w,null,{title:e(()=>[t(l(a(s)("data-planes.routes.item.mtls.supported_backends.title")),1)]),body:e(()=>[n("ul",null,[(o(!0),d(p,null,b(i.supportedBackends,u=>(o(),d("li",{key:u},l(u),1))),128))])]),_:2},1024)])]),_:2},1024))),128)):(o(),f(M,{key:1,class:"mt-4",appearance:"warning"},{alertMessage:e(()=>[n("div",{innerHTML:a(s)("data-planes.routes.item.mtls.disabled")},null,8,Lt)]),_:1}))]),t(),_.data.dataplaneInsight.subscriptions.length>0?(o(),d("div",Mt,[n("h2",null,l(a(s)("data-planes.routes.item.subscriptions.title")),1),t(),r(C,{class:"mt-4"},{default:e(()=>[r(mt,{subscriptions:_.data.dataplaneInsight.subscriptions},null,8,["subscriptions"])]),_:1})])):$("",!0)])]),_:2},[T.value.length>0||x?{name:"notifications",fn:e(()=>[n("ul",It,[(o(!0),d(p,null,b(T.value,i=>(o(),d("li",{key:i.kind,"data-testid":`warning-${i.kind}`,innerHTML:a(s)(`common.warnings.${i.kind}`,i.payload)},null,8,Dt))),128)),t(),x?(o(),d("li",Vt,[t(`
              The below view is not enhanced with runtime stats (Error loading stats: `),n("strong",null,l(x.toString()),1),t(`)
            `)])):$("",!0),t()])]),key:"0"}:void 0]),1024)]),_:2},1032,["src"])]),_:1})}}});const Jt=q(Ot,[["__scopeId","data-v-5a7bf611"]]);export{Jt as default};
