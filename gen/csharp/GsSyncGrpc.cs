// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: thetan/gateway/v1/gs_sync.proto
// </auto-generated>
#pragma warning disable 0414, 1591, 8981, 0612
#region Designer generated code

using grpc = global::Grpc.Core;

namespace Thetan.Gateway.V1 {
  public static partial class ThetanGateway
  {
    static readonly string __ServiceName = "thetan.gateway.v1.ThetanGateway";

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
    static readonly grpc::Marshaller<global::Thetan.Gateway.V1.PingRequest> __Marshaller_thetan_gateway_v1_PingRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Gateway.V1.PingRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.Gateway.V1.PingResponse> __Marshaller_thetan_gateway_v1_PingResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Gateway.V1.PingResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.Gateway.V1.PlayerConnectedRequest> __Marshaller_thetan_gateway_v1_PlayerConnectedRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Gateway.V1.PlayerConnectedRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.Gateway.V1.PlayerStatusResponse> __Marshaller_thetan_gateway_v1_PlayerStatusResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Gateway.V1.PlayerStatusResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.Gateway.V1.PlayerDisconnectedRequest> __Marshaller_thetan_gateway_v1_PlayerDisconnectedRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Gateway.V1.PlayerDisconnectedRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.Gateway.V1.RoomDestroyedRequest> __Marshaller_thetan_gateway_v1_RoomDestroyedRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Gateway.V1.RoomDestroyedRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.Gateway.V1.RoomDestroyedResponse> __Marshaller_thetan_gateway_v1_RoomDestroyedResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Gateway.V1.RoomDestroyedResponse.Parser));

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Thetan.Gateway.V1.PingRequest, global::Thetan.Gateway.V1.PingResponse> __Method_Ping = new grpc::Method<global::Thetan.Gateway.V1.PingRequest, global::Thetan.Gateway.V1.PingResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "Ping",
        __Marshaller_thetan_gateway_v1_PingRequest,
        __Marshaller_thetan_gateway_v1_PingResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Thetan.Gateway.V1.PlayerConnectedRequest, global::Thetan.Gateway.V1.PlayerStatusResponse> __Method_PlayerConnected = new grpc::Method<global::Thetan.Gateway.V1.PlayerConnectedRequest, global::Thetan.Gateway.V1.PlayerStatusResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "PlayerConnected",
        __Marshaller_thetan_gateway_v1_PlayerConnectedRequest,
        __Marshaller_thetan_gateway_v1_PlayerStatusResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Thetan.Gateway.V1.PlayerDisconnectedRequest, global::Thetan.Gateway.V1.PlayerStatusResponse> __Method_PlayerDisconnected = new grpc::Method<global::Thetan.Gateway.V1.PlayerDisconnectedRequest, global::Thetan.Gateway.V1.PlayerStatusResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "PlayerDisconnected",
        __Marshaller_thetan_gateway_v1_PlayerDisconnectedRequest,
        __Marshaller_thetan_gateway_v1_PlayerStatusResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Thetan.Gateway.V1.RoomDestroyedRequest, global::Thetan.Gateway.V1.RoomDestroyedResponse> __Method_RoomDestroyed = new grpc::Method<global::Thetan.Gateway.V1.RoomDestroyedRequest, global::Thetan.Gateway.V1.RoomDestroyedResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "RoomDestroyed",
        __Marshaller_thetan_gateway_v1_RoomDestroyedRequest,
        __Marshaller_thetan_gateway_v1_RoomDestroyedResponse);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
      get { return global::Thetan.Gateway.V1.GsSyncReflection.Descriptor.Services[0]; }
    }

    /// <summary>Base class for server-side implementations of ThetanGateway</summary>
    [grpc::BindServiceMethod(typeof(ThetanGateway), "BindService")]
    public abstract partial class ThetanGatewayBase
    {
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Thetan.Gateway.V1.PingResponse> Ping(global::Thetan.Gateway.V1.PingRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Thetan.Gateway.V1.PlayerStatusResponse> PlayerConnected(global::Thetan.Gateway.V1.PlayerConnectedRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Thetan.Gateway.V1.PlayerStatusResponse> PlayerDisconnected(global::Thetan.Gateway.V1.PlayerDisconnectedRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Thetan.Gateway.V1.RoomDestroyedResponse> RoomDestroyed(global::Thetan.Gateway.V1.RoomDestroyedRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

    }

    /// <summary>Client for ThetanGateway</summary>
    public partial class ThetanGatewayClient : grpc::ClientBase<ThetanGatewayClient>
    {
      /// <summary>Creates a new client for ThetanGateway</summary>
      /// <param name="channel">The channel to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public ThetanGatewayClient(grpc::ChannelBase channel) : base(channel)
      {
      }
      /// <summary>Creates a new client for ThetanGateway that uses a custom <c>CallInvoker</c>.</summary>
      /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public ThetanGatewayClient(grpc::CallInvoker callInvoker) : base(callInvoker)
      {
      }
      /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected ThetanGatewayClient() : base()
      {
      }
      /// <summary>Protected constructor to allow creation of configured clients.</summary>
      /// <param name="configuration">The client configuration.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected ThetanGatewayClient(ClientBaseConfiguration configuration) : base(configuration)
      {
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Gateway.V1.PingResponse Ping(global::Thetan.Gateway.V1.PingRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return Ping(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Gateway.V1.PingResponse Ping(global::Thetan.Gateway.V1.PingRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_Ping, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Gateway.V1.PingResponse> PingAsync(global::Thetan.Gateway.V1.PingRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return PingAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Gateway.V1.PingResponse> PingAsync(global::Thetan.Gateway.V1.PingRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_Ping, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Gateway.V1.PlayerStatusResponse PlayerConnected(global::Thetan.Gateway.V1.PlayerConnectedRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return PlayerConnected(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Gateway.V1.PlayerStatusResponse PlayerConnected(global::Thetan.Gateway.V1.PlayerConnectedRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_PlayerConnected, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Gateway.V1.PlayerStatusResponse> PlayerConnectedAsync(global::Thetan.Gateway.V1.PlayerConnectedRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return PlayerConnectedAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Gateway.V1.PlayerStatusResponse> PlayerConnectedAsync(global::Thetan.Gateway.V1.PlayerConnectedRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_PlayerConnected, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Gateway.V1.PlayerStatusResponse PlayerDisconnected(global::Thetan.Gateway.V1.PlayerDisconnectedRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return PlayerDisconnected(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Gateway.V1.PlayerStatusResponse PlayerDisconnected(global::Thetan.Gateway.V1.PlayerDisconnectedRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_PlayerDisconnected, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Gateway.V1.PlayerStatusResponse> PlayerDisconnectedAsync(global::Thetan.Gateway.V1.PlayerDisconnectedRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return PlayerDisconnectedAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Gateway.V1.PlayerStatusResponse> PlayerDisconnectedAsync(global::Thetan.Gateway.V1.PlayerDisconnectedRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_PlayerDisconnected, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Gateway.V1.RoomDestroyedResponse RoomDestroyed(global::Thetan.Gateway.V1.RoomDestroyedRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return RoomDestroyed(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Gateway.V1.RoomDestroyedResponse RoomDestroyed(global::Thetan.Gateway.V1.RoomDestroyedRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_RoomDestroyed, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Gateway.V1.RoomDestroyedResponse> RoomDestroyedAsync(global::Thetan.Gateway.V1.RoomDestroyedRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return RoomDestroyedAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Gateway.V1.RoomDestroyedResponse> RoomDestroyedAsync(global::Thetan.Gateway.V1.RoomDestroyedRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_RoomDestroyed, null, options, request);
      }
      /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected override ThetanGatewayClient NewInstance(ClientBaseConfiguration configuration)
      {
        return new ThetanGatewayClient(configuration);
      }
    }

    /// <summary>Creates service definition that can be registered with a server</summary>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static grpc::ServerServiceDefinition BindService(ThetanGatewayBase serviceImpl)
    {
      return grpc::ServerServiceDefinition.CreateBuilder()
          .AddMethod(__Method_Ping, serviceImpl.Ping)
          .AddMethod(__Method_PlayerConnected, serviceImpl.PlayerConnected)
          .AddMethod(__Method_PlayerDisconnected, serviceImpl.PlayerDisconnected)
          .AddMethod(__Method_RoomDestroyed, serviceImpl.RoomDestroyed).Build();
    }

    /// <summary>Register service method with a service binder with or without implementation. Useful when customizing the service binding logic.
    /// Note: this method is part of an experimental API that can change or be removed without any prior notice.</summary>
    /// <param name="serviceBinder">Service methods will be bound by calling <c>AddMethod</c> on this object.</param>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static void BindService(grpc::ServiceBinderBase serviceBinder, ThetanGatewayBase serviceImpl)
    {
      serviceBinder.AddMethod(__Method_Ping, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Thetan.Gateway.V1.PingRequest, global::Thetan.Gateway.V1.PingResponse>(serviceImpl.Ping));
      serviceBinder.AddMethod(__Method_PlayerConnected, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Thetan.Gateway.V1.PlayerConnectedRequest, global::Thetan.Gateway.V1.PlayerStatusResponse>(serviceImpl.PlayerConnected));
      serviceBinder.AddMethod(__Method_PlayerDisconnected, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Thetan.Gateway.V1.PlayerDisconnectedRequest, global::Thetan.Gateway.V1.PlayerStatusResponse>(serviceImpl.PlayerDisconnected));
      serviceBinder.AddMethod(__Method_RoomDestroyed, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Thetan.Gateway.V1.RoomDestroyedRequest, global::Thetan.Gateway.V1.RoomDestroyedResponse>(serviceImpl.RoomDestroyed));
    }

  }
}
#endregion
