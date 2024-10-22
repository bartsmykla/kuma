import{d as N,e as c,o,m as p,w as t,a,k as s,Y as d,b as e,t as i,a1 as A,p as f,c as n,H as m,J as g,l as S,a5 as K}from"./index-Bo5vSFZC.js";import{a as O,A as T}from"./AccordionList-D6aldBUn.js";import{C as j}from"./CodeBlock-BVhuix4S.js";import{P as F}from"./PolicyTypeTag-DkSTrlH5.js";import{T as I}from"./TagList-zfvB1RX9.js";import{R as M}from"./RuleMatchers-ztghSbVi.js";const Y={class:"stack-with-borders"},E={key:0,class:"mt-6"},H=s("h3",null,"Rules",-1),J={class:"mt-4"},q={class:"stack-with-borders"},G=s("dt",null,`
                                Config
                              `,-1),Q={class:"mt-2"},nt=N({__name:"ConnectionInboundSummaryOverviewView",props:{data:{}},setup(P){const r=P;return(U,W)=>{const C=c("KBadge"),B=c("RouterLink"),V=c("DataSource"),$=c("KCard"),x=c("DataCollection"),D=c("DataLoader"),L=c("AppView"),R=c("RouteView");return o(),p(R,{params:{mesh:"",dataPlane:"",connection:""},name:"connection-inbound-summary-overview-view"},{default:t(({t:b,route:v})=>[a(L,null,{default:t(()=>[s("div",Y,[a(d,{layout:"horizontal"},{title:t(()=>[e(`
            Tags
          `)]),body:t(()=>[a(I,{tags:r.data.tags,alignment:"right"},null,8,["tags"])]),_:1}),e(),a(d,{layout:"horizontal"},{title:t(()=>[e(i(b("http.api.property.state")),1)]),body:t(()=>[a(C,{appearance:r.data.state==="Ready"?"success":"danger"},{default:t(()=>[e(i(b(`http.api.value.${r.data.state}`)),1)]),_:2},1032,["appearance"])]),_:2},1024),e(),a(d,{layout:"horizontal"},{title:t(()=>[e(`
            Protocol
          `)]),body:t(()=>[a(C,{appearance:"info"},{default:t(()=>[e(i(b(`http.api.value.${r.data.protocol}`)),1)]),_:2},1024)]),_:2},1024),e(),a(d,{layout:"horizontal"},{title:t(()=>[e(`
            Address
          `)]),body:t(()=>[a(A,{text:`${r.data.addressPort}`},null,8,["text"])]),_:1}),e(),r.data.serviceAddressPort.length>0?(o(),p(d,{key:0,layout:"horizontal"},{title:t(()=>[e(`
            Service Address
          `)]),body:t(()=>[a(A,{text:`${r.data.serviceAddressPort}`},null,8,["text"])]),_:1})):f("",!0)]),e(),r.data?(o(),n("div",E,[H,e(),a(D,{src:`/meshes/${v.params.mesh}/rules/for/${v.params.dataPlane}`},{default:t(({data:z})=>[a(x,{predicate:h=>h.ruleType==="from"&&Number(h.inbound.port)===Number(v.params.connection.split("_")[1]),items:z.rules},{default:t(({items:h})=>[s("div",J,[a(O,{"initially-open":0,"multiple-open":"",class:"stack"},{default:t(()=>[(o(!0),n(m,null,g(Object.groupBy(h,y=>y.type),(y,k)=>(o(),p($,{key:k},{default:t(()=>[a(T,null,{"accordion-header":t(()=>[a(F,{"policy-type":k},{default:t(()=>[e(i(k)+" ("+i(y.length)+`)
                        `,1)]),_:2},1032,["policy-type"])]),"accordion-content":t(()=>[s("div",q,[(o(!0),n(m,null,g(y,u=>(o(),n(m,{key:u},[u.matchers.length>0?(o(),p(d,{key:0,layout:"horizontal"},{title:t(()=>[e(`
                                From
                              `)]),body:t(()=>[s("p",null,[a(M,{items:u.matchers},null,8,["items"])])]),_:2},1024)):f("",!0),e(),u.origins.length>0?(o(),p(d,{key:1,layout:"horizontal"},{title:t(()=>[e(`
                                Origin Policies
                              `)]),body:t(()=>[a(V,{src:"/policy-types"},{default:t(({data:w})=>[(o(!0),n(m,null,g([Object.groupBy((w==null?void 0:w.policies)??[],_=>_.name)],_=>(o(),n("ul",{key:_},[(o(!0),n(m,null,g(u.origins,l=>(o(),n("li",{key:`${l.mesh}-${l.name}`},[_[l.type]?(o(),p(B,{key:0,to:{name:"policy-detail-view",params:{mesh:l.mesh,policyPath:_[l.type][0].path,policy:l.name}}},{default:t(()=>[e(i(l.name),1)]),_:2},1032,["to"])):(o(),n(m,{key:1},[e(i(l.name),1)],64))]))),128))]))),128))]),_:2},1024)]),_:2},1024)):f("",!0),e(),s("div",null,[G,e(),s("dd",Q,[s("div",null,[a(j,{code:S(K).stringify(u.raw),language:"yaml","show-copy-button":!1},null,8,["code"])])])])],64))),128))])]),_:2},1024)]),_:2},1024))),128))]),_:2},1024)])]),_:2},1032,["predicate","items"])]),_:2},1032,["src"])])):f("",!0)]),_:2},1024)]),_:1})}}});export{nt as default};