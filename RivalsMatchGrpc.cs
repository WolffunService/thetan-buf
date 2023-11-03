// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: thetan/rivals/v1/rivals_match.proto
// </auto-generated>
#pragma warning disable 0414, 1591, 8981, 0612
#region Designer generated code

using grpc = global::Grpc.Core;

namespace Thetan.Rivals.V1 {
  public static partial class RivalMatchDirectorService
  {
    static readonly string __ServiceName = "thetan.rivals.v1.RivalMatchDirectorService";

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static void __Helper_SerializeMessage(global::Google.Protobuf.IMessage message, grpc::SerializationContext context)
    {
      #if !GRPC_DISABLE_PROTOBUF_BUFFER_SERIALIZATION
      if (message is global::Google.Protobuf.IBufferMessage)
      {
        context.SetPayloadLength(message.CalculateSize());
        global::Google.Protobuf.MessageExtensions.WriteTo(message, context.GetBufferWriter());
        context.Complete();
        return;
      }
      #endif
      context.Complete(global::Google.Protobuf.MessageExtensions.ToByteArray(message));
    }

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static class __Helper_MessageCache<T>
    {
      public static readonly bool IsBufferMessage = global::System.Reflection.IntrospectionExtensions.GetTypeInfo(typeof(global::Google.Protobuf.IBufferMessage)).IsAssignableFrom(typeof(T));
    }

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static T __Helper_DeserializeMessage<T>(grpc::DeserializationContext context, global::Google.Protobuf.MessageParser<T> parser) where T : global::Google.Protobuf.IMessage<T>
    {
      #if !GRPC_DISABLE_PROTOBUF_BUFFER_SERIALIZATION
      if (__Helper_MessageCache<T>.IsBufferMessage)
      {
        return parser.ParseFrom(context.PayloadAsReadOnlySequence());
      }
      #endif
      return parser.ParseFrom(context.PayloadAsNewBuffer());
    }

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.Rivals.V1.RivalCancelTicketRequest> __Marshaller_thetan_rivals_v1_RivalCancelTicketRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Rivals.V1.RivalCancelTicketRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.Rivals.V1.RivalCancelTicketResponse> __Marshaller_thetan_rivals_v1_RivalCancelTicketResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Rivals.V1.RivalCancelTicketResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.Rivals.V1.GetMatchInfoRequest> __Marshaller_thetan_rivals_v1_GetMatchInfoRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Rivals.V1.GetMatchInfoRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.Shared.V1.MatchFoundResponseProto> __Marshaller_thetan_shared_v1_MatchFoundResponseProto = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Shared.V1.MatchFoundResponseProto.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.Rivals.V1.CreateMatchOfflineRequest> __Marshaller_thetan_rivals_v1_CreateMatchOfflineRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Rivals.V1.CreateMatchOfflineRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.Rivals.V1.CreateMatchOfflineResponse> __Marshaller_thetan_rivals_v1_CreateMatchOfflineResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Rivals.V1.CreateMatchOfflineResponse.Parser));

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Thetan.Rivals.V1.RivalCancelTicketRequest, global::Thetan.Rivals.V1.RivalCancelTicketResponse> __Method_CancelTicket = new grpc::Method<global::Thetan.Rivals.V1.RivalCancelTicketRequest, global::Thetan.Rivals.V1.RivalCancelTicketResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "CancelTicket",
        __Marshaller_thetan_rivals_v1_RivalCancelTicketRequest,
        __Marshaller_thetan_rivals_v1_RivalCancelTicketResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Thetan.Rivals.V1.GetMatchInfoRequest, global::Thetan.Shared.V1.MatchFoundResponseProto> __Method_CreateMatchOnboard = new grpc::Method<global::Thetan.Rivals.V1.GetMatchInfoRequest, global::Thetan.Shared.V1.MatchFoundResponseProto>(
        grpc::MethodType.Unary,
        __ServiceName,
        "CreateMatchOnboard",
        __Marshaller_thetan_rivals_v1_GetMatchInfoRequest,
        __Marshaller_thetan_shared_v1_MatchFoundResponseProto);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Thetan.Rivals.V1.CreateMatchOfflineRequest, global::Thetan.Rivals.V1.CreateMatchOfflineResponse> __Method_CreateMatchOffline = new grpc::Method<global::Thetan.Rivals.V1.CreateMatchOfflineRequest, global::Thetan.Rivals.V1.CreateMatchOfflineResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "CreateMatchOffline",
        __Marshaller_thetan_rivals_v1_CreateMatchOfflineRequest,
        __Marshaller_thetan_rivals_v1_CreateMatchOfflineResponse);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
      get { return global::Thetan.Rivals.V1.RivalsMatchReflection.Descriptor.Services[0]; }
    }

    /// <summary>Base class for server-side implementations of RivalMatchDirectorService</summary>
    [grpc::BindServiceMethod(typeof(RivalMatchDirectorService), "BindService")]
    public abstract partial class RivalMatchDirectorServiceBase
    {
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Thetan.Rivals.V1.RivalCancelTicketResponse> CancelTicket(global::Thetan.Rivals.V1.RivalCancelTicketRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Thetan.Shared.V1.MatchFoundResponseProto> CreateMatchOnboard(global::Thetan.Rivals.V1.GetMatchInfoRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Thetan.Rivals.V1.CreateMatchOfflineResponse> CreateMatchOffline(global::Thetan.Rivals.V1.CreateMatchOfflineRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

    }

    /// <summary>Client for RivalMatchDirectorService</summary>
    public partial class RivalMatchDirectorServiceClient : grpc::ClientBase<RivalMatchDirectorServiceClient>
    {
      /// <summary>Creates a new client for RivalMatchDirectorService</summary>
      /// <param name="channel">The channel to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public RivalMatchDirectorServiceClient(grpc::ChannelBase channel) : base(channel)
      {
      }
      /// <summary>Creates a new client for RivalMatchDirectorService that uses a custom <c>CallInvoker</c>.</summary>
      /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public RivalMatchDirectorServiceClient(grpc::CallInvoker callInvoker) : base(callInvoker)
      {
      }
      /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected RivalMatchDirectorServiceClient() : base()
      {
      }
      /// <summary>Protected constructor to allow creation of configured clients.</summary>
      /// <param name="configuration">The client configuration.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected RivalMatchDirectorServiceClient(ClientBaseConfiguration configuration) : base(configuration)
      {
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Rivals.V1.RivalCancelTicketResponse CancelTicket(global::Thetan.Rivals.V1.RivalCancelTicketRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CancelTicket(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Rivals.V1.RivalCancelTicketResponse CancelTicket(global::Thetan.Rivals.V1.RivalCancelTicketRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_CancelTicket, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Rivals.V1.RivalCancelTicketResponse> CancelTicketAsync(global::Thetan.Rivals.V1.RivalCancelTicketRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CancelTicketAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Rivals.V1.RivalCancelTicketResponse> CancelTicketAsync(global::Thetan.Rivals.V1.RivalCancelTicketRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_CancelTicket, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Shared.V1.MatchFoundResponseProto CreateMatchOnboard(global::Thetan.Rivals.V1.GetMatchInfoRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CreateMatchOnboard(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Shared.V1.MatchFoundResponseProto CreateMatchOnboard(global::Thetan.Rivals.V1.GetMatchInfoRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_CreateMatchOnboard, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Shared.V1.MatchFoundResponseProto> CreateMatchOnboardAsync(global::Thetan.Rivals.V1.GetMatchInfoRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CreateMatchOnboardAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Shared.V1.MatchFoundResponseProto> CreateMatchOnboardAsync(global::Thetan.Rivals.V1.GetMatchInfoRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_CreateMatchOnboard, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Rivals.V1.CreateMatchOfflineResponse CreateMatchOffline(global::Thetan.Rivals.V1.CreateMatchOfflineRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CreateMatchOffline(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Rivals.V1.CreateMatchOfflineResponse CreateMatchOffline(global::Thetan.Rivals.V1.CreateMatchOfflineRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_CreateMatchOffline, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Rivals.V1.CreateMatchOfflineResponse> CreateMatchOfflineAsync(global::Thetan.Rivals.V1.CreateMatchOfflineRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CreateMatchOfflineAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Rivals.V1.CreateMatchOfflineResponse> CreateMatchOfflineAsync(global::Thetan.Rivals.V1.CreateMatchOfflineRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_CreateMatchOffline, null, options, request);
      }
      /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected override RivalMatchDirectorServiceClient NewInstance(ClientBaseConfiguration configuration)
      {
        return new RivalMatchDirectorServiceClient(configuration);
      }
    }

    /// <summary>Creates service definition that can be registered with a server</summary>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static grpc::ServerServiceDefinition BindService(RivalMatchDirectorServiceBase serviceImpl)
    {
      return grpc::ServerServiceDefinition.CreateBuilder()
          .AddMethod(__Method_CancelTicket, serviceImpl.CancelTicket)
          .AddMethod(__Method_CreateMatchOnboard, serviceImpl.CreateMatchOnboard)
          .AddMethod(__Method_CreateMatchOffline, serviceImpl.CreateMatchOffline).Build();
    }

    /// <summary>Register service method with a service binder with or without implementation. Useful when customizing the service binding logic.
    /// Note: this method is part of an experimental API that can change or be removed without any prior notice.</summary>
    /// <param name="serviceBinder">Service methods will be bound by calling <c>AddMethod</c> on this object.</param>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static void BindService(grpc::ServiceBinderBase serviceBinder, RivalMatchDirectorServiceBase serviceImpl)
    {
      serviceBinder.AddMethod(__Method_CancelTicket, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Thetan.Rivals.V1.RivalCancelTicketRequest, global::Thetan.Rivals.V1.RivalCancelTicketResponse>(serviceImpl.CancelTicket));
      serviceBinder.AddMethod(__Method_CreateMatchOnboard, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Thetan.Rivals.V1.GetMatchInfoRequest, global::Thetan.Shared.V1.MatchFoundResponseProto>(serviceImpl.CreateMatchOnboard));
      serviceBinder.AddMethod(__Method_CreateMatchOffline, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Thetan.Rivals.V1.CreateMatchOfflineRequest, global::Thetan.Rivals.V1.CreateMatchOfflineResponse>(serviceImpl.CreateMatchOffline));
    }

  }
}
#endregion
