// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: thetan/immortal/v1/immortal_bot.proto
// </auto-generated>
#pragma warning disable 0414, 1591, 8981, 0612
#region Designer generated code

using grpc = global::Grpc.Core;

namespace Thetan.Immortal.V1 {
  /// <summary>
  /// BotImmortalService is a grpc server for handling bot profile.
  /// This grpc is implemented in thetan-immortal-service and thetan-support calls it.
  /// </summary>
  public static partial class BotImmortalService
  {
    static readonly string __ServiceName = "thetan.immortal.v1.BotImmortalService";

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
    static readonly grpc::Marshaller<global::Thetan.Immortal.V1.SearchBotRankingRequest> __Marshaller_thetan_immortal_v1_SearchBotRankingRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Immortal.V1.SearchBotRankingRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.Immortal.V1.SearchBotRankingResponse> __Marshaller_thetan_immortal_v1_SearchBotRankingResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Immortal.V1.SearchBotRankingResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.Immortal.V1.CreateManyBotRankingRequest> __Marshaller_thetan_immortal_v1_CreateManyBotRankingRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Immortal.V1.CreateManyBotRankingRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.Immortal.V1.CreateManyBotRankingResponse> __Marshaller_thetan_immortal_v1_CreateManyBotRankingResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Immortal.V1.CreateManyBotRankingResponse.Parser));

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Thetan.Immortal.V1.SearchBotRankingRequest, global::Thetan.Immortal.V1.SearchBotRankingResponse> __Method_SearchBotRanking = new grpc::Method<global::Thetan.Immortal.V1.SearchBotRankingRequest, global::Thetan.Immortal.V1.SearchBotRankingResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "SearchBotRanking",
        __Marshaller_thetan_immortal_v1_SearchBotRankingRequest,
        __Marshaller_thetan_immortal_v1_SearchBotRankingResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Thetan.Immortal.V1.CreateManyBotRankingRequest, global::Thetan.Immortal.V1.CreateManyBotRankingResponse> __Method_CreateManyBotRanking = new grpc::Method<global::Thetan.Immortal.V1.CreateManyBotRankingRequest, global::Thetan.Immortal.V1.CreateManyBotRankingResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "CreateManyBotRanking",
        __Marshaller_thetan_immortal_v1_CreateManyBotRankingRequest,
        __Marshaller_thetan_immortal_v1_CreateManyBotRankingResponse);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
      get { return global::Thetan.Immortal.V1.ImmortalBotReflection.Descriptor.Services[0]; }
    }

    /// <summary>Base class for server-side implementations of BotImmortalService</summary>
    [grpc::BindServiceMethod(typeof(BotImmortalService), "BindService")]
    public abstract partial class BotImmortalServiceBase
    {
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Thetan.Immortal.V1.SearchBotRankingResponse> SearchBotRanking(global::Thetan.Immortal.V1.SearchBotRankingRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Thetan.Immortal.V1.CreateManyBotRankingResponse> CreateManyBotRanking(global::Thetan.Immortal.V1.CreateManyBotRankingRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

    }

    /// <summary>Client for BotImmortalService</summary>
    public partial class BotImmortalServiceClient : grpc::ClientBase<BotImmortalServiceClient>
    {
      /// <summary>Creates a new client for BotImmortalService</summary>
      /// <param name="channel">The channel to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public BotImmortalServiceClient(grpc::ChannelBase channel) : base(channel)
      {
      }
      /// <summary>Creates a new client for BotImmortalService that uses a custom <c>CallInvoker</c>.</summary>
      /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public BotImmortalServiceClient(grpc::CallInvoker callInvoker) : base(callInvoker)
      {
      }
      /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected BotImmortalServiceClient() : base()
      {
      }
      /// <summary>Protected constructor to allow creation of configured clients.</summary>
      /// <param name="configuration">The client configuration.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected BotImmortalServiceClient(ClientBaseConfiguration configuration) : base(configuration)
      {
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Immortal.V1.SearchBotRankingResponse SearchBotRanking(global::Thetan.Immortal.V1.SearchBotRankingRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return SearchBotRanking(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Immortal.V1.SearchBotRankingResponse SearchBotRanking(global::Thetan.Immortal.V1.SearchBotRankingRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_SearchBotRanking, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Immortal.V1.SearchBotRankingResponse> SearchBotRankingAsync(global::Thetan.Immortal.V1.SearchBotRankingRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return SearchBotRankingAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Immortal.V1.SearchBotRankingResponse> SearchBotRankingAsync(global::Thetan.Immortal.V1.SearchBotRankingRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_SearchBotRanking, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Immortal.V1.CreateManyBotRankingResponse CreateManyBotRanking(global::Thetan.Immortal.V1.CreateManyBotRankingRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CreateManyBotRanking(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Immortal.V1.CreateManyBotRankingResponse CreateManyBotRanking(global::Thetan.Immortal.V1.CreateManyBotRankingRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_CreateManyBotRanking, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Immortal.V1.CreateManyBotRankingResponse> CreateManyBotRankingAsync(global::Thetan.Immortal.V1.CreateManyBotRankingRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CreateManyBotRankingAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Immortal.V1.CreateManyBotRankingResponse> CreateManyBotRankingAsync(global::Thetan.Immortal.V1.CreateManyBotRankingRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_CreateManyBotRanking, null, options, request);
      }
      /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected override BotImmortalServiceClient NewInstance(ClientBaseConfiguration configuration)
      {
        return new BotImmortalServiceClient(configuration);
      }
    }

    /// <summary>Creates service definition that can be registered with a server</summary>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static grpc::ServerServiceDefinition BindService(BotImmortalServiceBase serviceImpl)
    {
      return grpc::ServerServiceDefinition.CreateBuilder()
          .AddMethod(__Method_SearchBotRanking, serviceImpl.SearchBotRanking)
          .AddMethod(__Method_CreateManyBotRanking, serviceImpl.CreateManyBotRanking).Build();
    }

    /// <summary>Register service method with a service binder with or without implementation. Useful when customizing the service binding logic.
    /// Note: this method is part of an experimental API that can change or be removed without any prior notice.</summary>
    /// <param name="serviceBinder">Service methods will be bound by calling <c>AddMethod</c> on this object.</param>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static void BindService(grpc::ServiceBinderBase serviceBinder, BotImmortalServiceBase serviceImpl)
    {
      serviceBinder.AddMethod(__Method_SearchBotRanking, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Thetan.Immortal.V1.SearchBotRankingRequest, global::Thetan.Immortal.V1.SearchBotRankingResponse>(serviceImpl.SearchBotRanking));
      serviceBinder.AddMethod(__Method_CreateManyBotRanking, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Thetan.Immortal.V1.CreateManyBotRankingRequest, global::Thetan.Immortal.V1.CreateManyBotRankingResponse>(serviceImpl.CreateManyBotRanking));
    }

  }
}
#endregion
