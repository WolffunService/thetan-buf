// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: thetan/bot/v1/bot_rivals.proto
// </auto-generated>
#pragma warning disable 0414, 1591, 8981, 0612
#region Designer generated code

using grpc = global::Grpc.Core;

namespace Thetan.Rivals.V1 {
  public static partial class BotRivalsService
  {
    static readonly string __ServiceName = "thetan.rivals.v1.BotRivalsService";

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
    static readonly grpc::Marshaller<global::Thetan.Rivals.V1.FetchLobbyBotsRequest> __Marshaller_thetan_rivals_v1_FetchLobbyBotsRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Rivals.V1.FetchLobbyBotsRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.Rivals.V1.FetchLobbyBotsResponse> __Marshaller_thetan_rivals_v1_FetchLobbyBotsResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Rivals.V1.FetchLobbyBotsResponse.Parser));

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Thetan.Rivals.V1.FetchLobbyBotsRequest, global::Thetan.Rivals.V1.FetchLobbyBotsResponse> __Method_FetchLobbyBots = new grpc::Method<global::Thetan.Rivals.V1.FetchLobbyBotsRequest, global::Thetan.Rivals.V1.FetchLobbyBotsResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "FetchLobbyBots",
        __Marshaller_thetan_rivals_v1_FetchLobbyBotsRequest,
        __Marshaller_thetan_rivals_v1_FetchLobbyBotsResponse);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
      get { return global::Thetan.Rivals.V1.BotRivalsReflection.Descriptor.Services[0]; }
    }

    /// <summary>Base class for server-side implementations of BotRivalsService</summary>
    [grpc::BindServiceMethod(typeof(BotRivalsService), "BindService")]
    public abstract partial class BotRivalsServiceBase
    {
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Thetan.Rivals.V1.FetchLobbyBotsResponse> FetchLobbyBots(global::Thetan.Rivals.V1.FetchLobbyBotsRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

    }

    /// <summary>Client for BotRivalsService</summary>
    public partial class BotRivalsServiceClient : grpc::ClientBase<BotRivalsServiceClient>
    {
      /// <summary>Creates a new client for BotRivalsService</summary>
      /// <param name="channel">The channel to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public BotRivalsServiceClient(grpc::ChannelBase channel) : base(channel)
      {
      }
      /// <summary>Creates a new client for BotRivalsService that uses a custom <c>CallInvoker</c>.</summary>
      /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public BotRivalsServiceClient(grpc::CallInvoker callInvoker) : base(callInvoker)
      {
      }
      /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected BotRivalsServiceClient() : base()
      {
      }
      /// <summary>Protected constructor to allow creation of configured clients.</summary>
      /// <param name="configuration">The client configuration.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected BotRivalsServiceClient(ClientBaseConfiguration configuration) : base(configuration)
      {
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Rivals.V1.FetchLobbyBotsResponse FetchLobbyBots(global::Thetan.Rivals.V1.FetchLobbyBotsRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return FetchLobbyBots(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Rivals.V1.FetchLobbyBotsResponse FetchLobbyBots(global::Thetan.Rivals.V1.FetchLobbyBotsRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_FetchLobbyBots, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Rivals.V1.FetchLobbyBotsResponse> FetchLobbyBotsAsync(global::Thetan.Rivals.V1.FetchLobbyBotsRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return FetchLobbyBotsAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Rivals.V1.FetchLobbyBotsResponse> FetchLobbyBotsAsync(global::Thetan.Rivals.V1.FetchLobbyBotsRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_FetchLobbyBots, null, options, request);
      }
      /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected override BotRivalsServiceClient NewInstance(ClientBaseConfiguration configuration)
      {
        return new BotRivalsServiceClient(configuration);
      }
    }

    /// <summary>Creates service definition that can be registered with a server</summary>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static grpc::ServerServiceDefinition BindService(BotRivalsServiceBase serviceImpl)
    {
      return grpc::ServerServiceDefinition.CreateBuilder()
          .AddMethod(__Method_FetchLobbyBots, serviceImpl.FetchLobbyBots).Build();
    }

    /// <summary>Register service method with a service binder with or without implementation. Useful when customizing the service binding logic.
    /// Note: this method is part of an experimental API that can change or be removed without any prior notice.</summary>
    /// <param name="serviceBinder">Service methods will be bound by calling <c>AddMethod</c> on this object.</param>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static void BindService(grpc::ServiceBinderBase serviceBinder, BotRivalsServiceBase serviceImpl)
    {
      serviceBinder.AddMethod(__Method_FetchLobbyBots, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Thetan.Rivals.V1.FetchLobbyBotsRequest, global::Thetan.Rivals.V1.FetchLobbyBotsResponse>(serviceImpl.FetchLobbyBots));
    }

  }
}
#endregion
