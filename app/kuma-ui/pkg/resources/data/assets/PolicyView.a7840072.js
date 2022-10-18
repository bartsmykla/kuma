import{d as $,o,c as C,w as s,a as y,b as S,r as E,e as M,K as N,f as w,g as e,h as Q,v as X,F as H,i as O,t as I,u as ee,j as r,k as F,l as te,P as z,m as R,s as U,n as h,p as D,q as ae,x as se,y as le,z as ne}from"./index.d0c2efa7.js";import{D as oe}from"./DataOverview.c3903863.js";import{E as re}from"./EntityURLControl.b14a7a6c.js";import{F as ie}from"./FrameSkeleton.fba0c2e6.js";import{L as W}from"./LabelList.84abaca0.js";import{T as ue}from"./TabsWidget.d742b9a7.js";import{Y as ce}from"./YamlView.2cfc9d9d.js";import"./EmptyBlock.vue_vue_type_script_setup_true_lang.37c6fb15.js";import"./ErrorBlock.0c813134.js";import"./LoadingBlock.vue_vue_type_script_setup_true_lang.d2d6c4b6.js";import"./index.58caa11d.js";import"./CodeBlock.22897d3f.js";import"./_commonjsHelpers.f037b798.js";const pe=$({__name:"DocumentationLink",props:{href:{type:String,required:!0}},setup(a){return(i,q)=>{const f=E("KIcon"),u=E("KButton");return o(),C(u,{class:"docs-link",appearance:"outline",target:"_blank",to:a.href},{icon:s(()=>[y(f,{icon:"externalLink"})]),default:s(()=>[S(" Documentation ")]),_:1},8,["to"])}}}),me={name:"PolicyConnections",components:{LabelList:W},props:{mesh:{type:String,required:!0},policyType:{type:String,required:!0},policyName:{type:String,required:!0}},data(){return{hasDataplanes:!1,isLoading:!0,hasError:!1,dataplanes:[],searchInput:""}},computed:{filteredDataplanes(){const a=this.searchInput.toLowerCase();return this.dataplanes.filter(({dataplane:{name:i}})=>i.toLowerCase().includes(a))}},watch:{policyName(){this.fetchPolicyConntections()}},mounted(){this.fetchPolicyConntections()},methods:{async fetchPolicyConntections(){this.hasError=!1,this.isLoading=!0;try{const{items:a,total:i}=await N.getPolicyConnections({mesh:this.mesh,policyType:this.policyType,policyName:this.policyName});this.hasDataplanes=i>0,this.dataplanes=a}catch{this.hasError=!0}finally{this.isLoading=!1}}}},de=e("h4",null,"Dataplanes",-1);function he(a,i,q,f,u,_){const g=E("router-link"),p=E("LabelList");return o(),w("div",null,[y(p,{"has-error":u.hasError,"is-loading":u.isLoading,"is-empty":!u.hasDataplanes},{default:s(()=>[e("ul",null,[e("li",null,[de,Q(e("input",{id:"dataplane-search","onUpdate:modelValue":i[0]||(i[0]=l=>u.searchInput=l),type:"text",class:"k-input mb-4",placeholder:"Filter by name",required:""},null,512),[[X,u.searchInput]]),(o(!0),w(H,null,O(_.filteredDataplanes,(l,v)=>(o(),w("p",{key:v,class:"my-1","data-testid":"dataplane-name"},[y(g,{to:{name:"data-plane-list-view",query:{ns:l.dataplane.name},params:{mesh:l.dataplane.mesh}}},{default:s(()=>[S(I(l.dataplane.name),1)]),_:2},1032,["to"])]))),128))])])]),_:1},8,["has-error","is-loading","is-empty"])])}const ye=M(me,[["render",he]]),Y=a=>(le("data-v-f90da74a"),a=a(),ne(),a),ve={key:0,class:"mb-4"},fe=Y(()=>e("p",null,[e("strong",null,"Warning"),S(" This policy is experimental. If you encountered any problem please open an "),e("a",{href:"https://github.com/kumahq/kuma/issues/new/choose",target:"_blank",rel:"noopener noreferrer"},"issue")],-1)),_e=Y(()=>e("span",{class:"custom-control-icon"}," \u2190 ",-1)),ge={"data-testid":"policy-single-entity"},be={"data-testid":"policy-overview-tab"},ke={class:"config-wrapper"},we=$({__name:"PolicyView",props:{policyPath:{type:String,required:!0}},setup(a){const i=a,q=[{hash:"#overview",title:"Overview"},{hash:"#affected-dpps",title:"Affected DPPs"}],f=ee(),u=se(),_=r(!0),g=r(!1),p=r(null),l=r(!0),v=r(!1),L=r(!1),T=r(!1),P=r({}),b=r(null),K=r(null),A=r({headers:[{label:"Actions",key:"actions",hideLabel:!0},{label:"Name",key:"name"},{label:"Mesh",key:"mesh"},{label:"Type",key:"type"}],data:[]}),c=F(()=>u.state.policiesByPath[i.policyPath]),j=F(()=>`https://kuma.io/docs/${u.getters["config/getKumaDocsVersion"]}/policies/${c.value.path}/`);te(()=>f.params.mesh,function(){f.name===i.policyPath&&(_.value=!0,g.value=!1,l.value=!0,v.value=!1,L.value=!1,T.value=!1,p.value=null,B())}),B();async function B(t=0){_.value=!0,p.value=null;const n=f.query.ns||null,m=f.params.mesh,V=c.value.path;try{let d;if(m!==null&&n!==null)d=[await N.getSinglePolicyEntity({mesh:m,path:V,name:n})],K.value=null;else{const k={size:z,offset:t},x=await N.getAllPolicyEntitiesFromMesh({mesh:m,path:V},k);d=x.items,K.value=x.next}if(d.length>0){A.value.data=d.map(J=>G(J)),T.value=!1,g.value=!1;const k=["type","name","mesh"],x=d[0];P.value=R(x,k),b.value=U(x)}else A.value.data=[],T.value=!0,g.value=!0,v.value=!0}catch(d){p.value=d,g.value=!0}finally{_.value=!1,l.value=!1}}function G(t){if(!t.mesh)return t;const n=t,m={name:"mesh-detail-view",params:{mesh:t.mesh}};return n.meshRoute=m,n}async function Z(t){if(L.value=!1,l.value=!0,v.value=!1,t)try{const n=await N.getSinglePolicyEntity({mesh:t.mesh,path:c.value.path,name:t.name});if(n){const m=["type","name","mesh"];t.value=R(n,m),b.value=U(n)}else t.value={},v.value=!0}catch(n){L.value=!0,console.error(n)}finally{l.value=!1}}return(t,n)=>{const m=E("KAlert"),V=E("KButton");return h(c)?(o(),w("div",{key:0,class:ae(["relative",h(c).path])},[h(c).isExperimental?(o(),w("div",ve,[y(m,{appearance:"warning"},{alertMessage:s(()=>[fe]),_:1})])):D("",!0),y(ie,null,{default:s(()=>[y(oe,{"page-size":h(z),"has-error":p.value!==null,error:p.value,"is-loading":_.value,"empty-state":{title:"No Data",message:`There are no ${h(c).pluralDisplayName} present.`},"table-data":A.value,"table-data-is-empty":T.value,next:K.value,onTableAction:Z,onLoadData:B},{additionalControls:s(()=>[y(pe,{href:h(j),"data-testid":"policy-documentation-link"},null,8,["href"]),t.$route.query.ns?(o(),C(V,{key:0,class:"back-button",appearance:"primary",to:{name:h(c).path}},{default:s(()=>[_e,S(" View All ")]),_:1},8,["to"])):D("",!0)]),default:s(()=>[S(" > ")]),_:1},8,["page-size","has-error","error","is-loading","empty-state","table-data","table-data-is-empty","next"]),g.value===!1?(o(),C(ue,{key:0,"has-error":p.value!==null,error:p.value,"is-loading":_.value,tabs:q,"initial-tab-override":"overview"},{tabHeader:s(()=>[e("div",null,[e("h3",ge,I(h(c).singularDisplayName)+": "+I(P.value.name),1)]),e("div",null,[y(re,{name:P.value.name,mesh:P.value.mesh},null,8,["name","mesh"])])]),overview:s(()=>[y(W,{"has-error":L.value,"is-loading":l.value,"is-empty":v.value},{default:s(()=>[e("div",be,[e("ul",null,[(o(!0),w(H,null,O(P.value,(d,k)=>(o(),w("li",{key:k},[e("h4",null,I(k),1),e("p",null,I(d),1)]))),128))])])]),_:1},8,["has-error","is-loading","is-empty"]),e("div",ke,[b.value!==null?(o(),C(ce,{key:0,id:"code-block-policy","has-error":L.value,"is-loading":l.value,"is-empty":v.value,content:b.value,"is-searchable":""},null,8,["has-error","is-loading","is-empty","content"])):D("",!0)])]),"affected-dpps":s(()=>[b.value!==null?(o(),C(ye,{key:0,mesh:b.value.mesh,"policy-name":b.value.name,"policy-type":h(c).path},null,8,["mesh","policy-name","policy-type"])):D("",!0)]),_:1},8,["has-error","error","is-loading"])):D("",!0)]),_:1})],2)):D("",!0)}}});const Ae=M(we,[["__scopeId","data-v-f90da74a"]]);export{Ae as default};
