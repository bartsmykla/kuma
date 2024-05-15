import{a as N,A as S}from"./AccordionList-CWZyY1fF.js";import{C as T}from"./CodeBlock-C_tPei7l.js";import{d as K,a as c,o,b as u,w as t,e as a,m as s,a7 as d,f as e,t as m,T as C,q as f,c as n,F as p,G as b,p as O}from"./index-DIXzft6Y.js";import{P as F}from"./PolicyTypeTag-Do6a-zE_.js";import{T as j}from"./TagList-BlSwqgEl.js";import{R as I}from"./RuleMatchers-BHRh9DXC.js";import{t as q}from"./toYaml-DB9FPXFY.js";const E={class:"stack-with-borders"},G={key:0,class:"mt-6"},H=s("h3",null,"Rules",-1),M={class:"mt-4"},U={class:"stack-with-borders"},W=s("dt",null,`
                                Config
                              `,-1),Y={class:"mt-2"},st=K({__name:"ConnectionInboundSummaryOverviewView",props:{data:{}},setup(B){const r=B;return(J,Q)=>{const w=c("KBadge"),P=c("RouterLink"),x=c("DataSource"),A=c("KCard"),V=c("DataCollection"),D=c("DataLoader"),z=c("AppView"),L=c("RouteView");return o(),u(L,{params:{mesh:"",dataPlane:"",connection:""},name:"connection-inbound-summary-overview-view"},{default:t(({t:R,route:g})=>[a(z,null,{default:t(()=>[s("div",E,[a(d,{layout:"horizontal"},{title:t(()=>[e(`
            Tags
          `)]),body:t(()=>[a(j,{tags:r.data.tags,alignment:"right"},null,8,["tags"])]),_:1}),e(),a(d,{layout:"horizontal"},{title:t(()=>[e(`
            Status
          `)]),body:t(()=>[a(w,{appearance:r.data.health.ready?"success":"danger"},{default:t(()=>[e(m(r.data.health.ready?"Healthy":"Unhealthy"),1)]),_:1},8,["appearance"])]),_:1}),e(),a(d,{layout:"horizontal"},{title:t(()=>[e(`
            Protocol
          `)]),body:t(()=>[a(w,{appearance:"info"},{default:t(()=>[e(m(R(`http.api.value.${r.data.protocol}`)),1)]),_:2},1024)]),_:2},1024),e(),a(d,{layout:"horizontal"},{title:t(()=>[e(`
            Address
          `)]),body:t(()=>[a(C,{text:`${r.data.addressPort}`},null,8,["text"])]),_:1}),e(),r.data.serviceAddressPort.length>0?(o(),u(d,{key:0,layout:"horizontal"},{title:t(()=>[e(`
            Service Address
          `)]),body:t(()=>[a(C,{text:`${r.data.serviceAddressPort}`},null,8,["text"])]),_:1})):f("",!0)]),e(),r.data?(o(),n("div",G,[H,e(),a(D,{src:`/meshes/${g.params.mesh}/rules/for/${g.params.dataPlane}`},{default:t(({data:$})=>[a(V,{predicate:h=>h.ruleType==="from"&&Number(h.inbound.port)===Number(g.params.connection.split("_")[1]),items:$.rules},{default:t(({items:h})=>[s("div",M,[a(N,{"initially-open":0,"multiple-open":"",class:"stack"},{default:t(()=>[(o(!0),n(p,null,b(Object.groupBy(h,y=>y.type),(y,v)=>(o(),u(A,{key:v},{default:t(()=>[a(S,null,{"accordion-header":t(()=>[a(F,{"policy-type":v},{default:t(()=>[e(m(v)+" ("+m(y.length)+`)
                        `,1)]),_:2},1032,["policy-type"])]),"accordion-content":t(()=>[s("div",U,[(o(!0),n(p,null,b(y,i=>(o(),n(p,{key:i},[i.matchers.length>0?(o(),u(d,{key:0,layout:"horizontal"},{title:t(()=>[e(`
                                From
                              `)]),body:t(()=>[s("p",null,[a(I,{items:i.matchers},null,8,["items"])])]),_:2},1024)):f("",!0),e(),i.origins.length>0?(o(),u(d,{key:1,layout:"horizontal"},{title:t(()=>[e(`
                                Origin Policies
                              `)]),body:t(()=>[a(x,{src:"/policy-types"},{default:t(({data:k})=>[(o(!0),n(p,null,b([Object.groupBy((k==null?void 0:k.policies)??[],_=>_.name)],_=>(o(),n("ul",{key:_},[(o(!0),n(p,null,b(i.origins,l=>(o(),n("li",{key:`${l.mesh}-${l.name}`},[_[l.type]?(o(),u(P,{key:0,to:{name:"policy-detail-view",params:{mesh:l.mesh,policyPath:_[l.type][0].path,policy:l.name}}},{default:t(()=>[e(m(l.name),1)]),_:2},1032,["to"])):(o(),n(p,{key:1},[e(m(l.name),1)],64))]))),128))]))),128))]),_:2},1024)]),_:2},1024)):f("",!0),e(),s("div",null,[W,e(),s("dd",Y,[s("div",null,[a(T,{code:O(q)(i.raw),language:"yaml","show-copy-button":!1},null,8,["code"])])])])],64))),128))])]),_:2},1024)]),_:2},1024))),128))]),_:2},1024)])]),_:2},1032,["predicate","items"])]),_:2},1032,["src"])])):f("",!0)]),_:2},1024)]),_:1})}}});export{st as default};