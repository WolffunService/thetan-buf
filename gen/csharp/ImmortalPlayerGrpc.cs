// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: thetan/immortal/v1/immortal_player.proto
// </auto-generated>
#pragma warning disable 0414, 1591, 8981, 0612
#region Designer generated code

using grpc = global::Grpc.Core;

namespace Thetan.Immortal.V1 {
  /// <summary>
  /// ImmortalPlayerService is a grpc server for handling player profile.
  /// This grpc is implemented in thetan-player-immortal and match-director-immortal calls it.
  /// </summary>
  public static partial class ImmortalPlayerService
  {
    static readonly string __ServiceName = "thetan.immortal.v1.ImmortalPlayerService";

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
    static readonly grpc::Marshaller<global::Thetan.Immortal.V1.SearchPlayersRequest> __Marshaller_thetan_immortal_v1_SearchPlayersRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Immortal.V1.SearchPlayersRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.Immortal.V1.SearchPlayersResponse> __Marshaller_thetan_immortal_v1_SearchPlayersResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Immortal.V1.SearchPlayersResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.Immortal.V1.GetBotNamesRequest> __Marshaller_thetan_immortal_v1_GetBotNamesRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Immortal.V1.GetBotNamesRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.Immortal.V1.GetBotNamesResponse> __Marshaller_thetan_immortal_v1_GetBotNamesResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Immortal.V1.GetBotNamesResponse.Parser));

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Thetan.Immortal.V1.SearchPlayersRequest, global::Thetan.Immortal.V1.SearchPlayersResponse> __Method_SearchPlayers = new grpc::Method<global::Thetan.Immortal.V1.SearchPlayersRequest, global::Thetan.Immortal.V1.SearchPlayersResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "SearchPlayers",
        __Marshaller_thetan_immortal_v1_SearchPlayersRequest,
        __Marshaller_thetan_immortal_v1_SearchPlayersResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Thetan.Immortal.V1.GetBotNamesRequest, global::Thetan.Immortal.V1.GetBotNamesResponse> __Method_GetBotNames = new grpc::Method<global::Thetan.Immortal.V1.GetBotNamesRequest, global::Thetan.Immortal.V1.GetBotNamesResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "GetBotNames",
        __Marshaller_thetan_immortal_v1_GetBotNamesRequest,
        __Marshaller_thetan_immortal_v1_GetBotNamesResponse);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
      get { return global::Thetan.Immortal.V1.ImmortalPlayerReflection.Descriptor.Services[0]; }
    }

    /// <summary>Base class for server-side implementations of ImmortalPlayerService</summary>
    [grpc::BindServiceMethod(typeof(ImmortalPlayerService), "BindService")]
    public abstract partial class ImmortalPlayerServiceBase
    {
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Thetan.Immortal.V1.SearchPlayersResponse> SearchPlayers(global::Thetan.Immortal.V1.SearchPlayersRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Thetan.Immortal.V1.GetBotNamesResponse> GetBotNames(global::Thetan.Immortal.V1.GetBotNamesRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

    }

    /// <summary>Client for ImmortalPlayerService</summary>
    public partial class ImmortalPlayerServiceClient : grpc::ClientBase<ImmortalPlayerServiceClient>
    {
      /// <summary>Creates a new client for ImmortalPlayerService</summary>
      /// <param name="channel">The channel to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public ImmortalPlayerServiceClient(grpc::ChannelBase channel) : base(channel)
      {
      }
      /// <summary>Creates a new client for ImmortalPlayerService that uses a custom <c>CallInvoker</c>.</summary>
      /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public ImmortalPlayerServiceClient(grpc::CallInvoker callInvoker) : base(callInvoker)
      {
      }
      /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected ImmortalPlayerServiceClient() : base()
      {
      }
      /// <summary>Protected constructor to allow creation of configured clients.</summary>
      /// <param name="configuration">The client configuration.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected ImmortalPlayerServiceClient(ClientBaseConfiguration configuration) : base(configuration)
      {
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Immortal.V1.SearchPlayersResponse SearchPlayers(global::Thetan.Immortal.V1.SearchPlayersRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return SearchPlayers(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Immortal.V1.SearchPlayersResponse SearchPlayers(global::Thetan.Immortal.V1.SearchPlayersRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_SearchPlayers, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Immortal.V1.SearchPlayersResponse> SearchPlayersAsync(global::Thetan.Immortal.V1.SearchPlayersRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return SearchPlayersAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Immortal.V1.SearchPlayersResponse> SearchPlayersAsync(global::Thetan.Immortal.V1.SearchPlayersRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_SearchPlayers, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Immortal.V1.GetBotNamesResponse GetBotNames(global::Thetan.Immortal.V1.GetBotNamesRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetBotNames(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Immortal.V1.GetBotNamesResponse GetBotNames(global::Thetan.Immortal.V1.GetBotNamesRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_GetBotNames, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Immortal.V1.GetBotNamesResponse> GetBotNamesAsync(global::Thetan.Immortal.V1.GetBotNamesRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetBotNamesAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Immortal.V1.GetBotNamesResponse> GetBotNamesAsync(global::Thetan.Immortal.V1.GetBotNamesRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_GetBotNames, null, options, request);
      }
      /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected override ImmortalPlayerServiceClient NewInstance(ClientBaseConfiguration configuration)
      {
        return new ImmortalPlayerServiceClient(configuration);
      }
    }

    /// <summary>Creates service definition that can be registered with a server</summary>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static grpc::ServerServiceDefinition BindService(ImmortalPlayerServiceBase serviceImpl)
    {
      return grpc::ServerServiceDefinition.CreateBuilder()
          .AddMethod(__Method_SearchPlayers, serviceImpl.SearchPlayers)
          .AddMethod(__Method_GetBotNames, serviceImpl.GetBotNames).Build();
    }

    /// <summary>Register service method with a service binder with or without implementation. Useful when customizing the service binding logic.
    /// Note: this method is part of an experimental API that can change or be removed without any prior notice.</summary>
    /// <param name="serviceBinder">Service methods will be bound by calling <c>AddMethod</c> on this object.</param>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static void BindService(grpc::ServiceBinderBase serviceBinder, ImmortalPlayerServiceBase serviceImpl)
    {
      serviceBinder.AddMethod(__Method_SearchPlayers, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Thetan.Immortal.V1.SearchPlayersRequest, global::Thetan.Immortal.V1.SearchPlayersResponse>(serviceImpl.SearchPlayers));
      serviceBinder.AddMethod(__Method_GetBotNames, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Thetan.Immortal.V1.GetBotNamesRequest, global::Thetan.Immortal.V1.GetBotNamesResponse>(serviceImpl.GetBotNames));
    }

  }
}
#endregion
