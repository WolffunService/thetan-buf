// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: thetan/shared/v1/match.proto
// </auto-generated>
#pragma warning disable 0414, 1591, 8981, 0612
#region Designer generated code

using grpc = global::Grpc.Core;

namespace Thetan.Shared.V1 {
  /// <summary>
  /// FE call request match
  /// </summary>
  public static partial class MatchService
  {
    static readonly string __ServiceName = "thetan.shared.v1.MatchService";

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
    static readonly grpc::Marshaller<global::Thetan.Shared.V1.MatchProtoVersionPackage> __Marshaller_thetan_shared_v1_MatchProtoVersionPackage = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Shared.V1.MatchProtoVersionPackage.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.Shared.V1.MatchProtoPackage> __Marshaller_thetan_shared_v1_MatchProtoPackage = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Shared.V1.MatchProtoPackage.Parser));

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Thetan.Shared.V1.MatchProtoVersionPackage, global::Thetan.Shared.V1.MatchProtoPackage> __Method_CreateMatchWithVersion = new grpc::Method<global::Thetan.Shared.V1.MatchProtoVersionPackage, global::Thetan.Shared.V1.MatchProtoPackage>(
        grpc::MethodType.ServerStreaming,
        __ServiceName,
        "CreateMatchWithVersion",
        __Marshaller_thetan_shared_v1_MatchProtoVersionPackage,
        __Marshaller_thetan_shared_v1_MatchProtoPackage);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Thetan.Shared.V1.MatchProtoPackage, global::Thetan.Shared.V1.MatchProtoPackage> __Method_RegisterMatchFound = new grpc::Method<global::Thetan.Shared.V1.MatchProtoPackage, global::Thetan.Shared.V1.MatchProtoPackage>(
        grpc::MethodType.ServerStreaming,
        __ServiceName,
        "RegisterMatchFound",
        __Marshaller_thetan_shared_v1_MatchProtoPackage,
        __Marshaller_thetan_shared_v1_MatchProtoPackage);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Thetan.Shared.V1.MatchProtoPackage, global::Thetan.Shared.V1.MatchProtoPackage> __Method_CancelMatchMaking = new grpc::Method<global::Thetan.Shared.V1.MatchProtoPackage, global::Thetan.Shared.V1.MatchProtoPackage>(
        grpc::MethodType.Unary,
        __ServiceName,
        "CancelMatchMaking",
        __Marshaller_thetan_shared_v1_MatchProtoPackage,
        __Marshaller_thetan_shared_v1_MatchProtoPackage);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
      get { return global::Thetan.Shared.V1.MatchReflection.Descriptor.Services[0]; }
    }

    /// <summary>Base class for server-side implementations of MatchService</summary>
    [grpc::BindServiceMethod(typeof(MatchService), "BindService")]
    public abstract partial class MatchServiceBase
    {
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task CreateMatchWithVersion(global::Thetan.Shared.V1.MatchProtoVersionPackage request, grpc::IServerStreamWriter<global::Thetan.Shared.V1.MatchProtoPackage> responseStream, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      /// <summary>
      ///register de nhan su kien match found
      /// </summary>
      /// <param name="request">The request received from the client.</param>
      /// <param name="responseStream">Used for sending responses back to the client.</param>
      /// <param name="context">The context of the server-side call handler being invoked.</param>
      /// <returns>A task indicating completion of the handler.</returns>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task RegisterMatchFound(global::Thetan.Shared.V1.MatchProtoPackage request, grpc::IServerStreamWriter<global::Thetan.Shared.V1.MatchProtoPackage> responseStream, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Thetan.Shared.V1.MatchProtoPackage> CancelMatchMaking(global::Thetan.Shared.V1.MatchProtoPackage request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

    }

    /// <summary>Client for MatchService</summary>
    public partial class MatchServiceClient : grpc::ClientBase<MatchServiceClient>
    {
      /// <summary>Creates a new client for MatchService</summary>
      /// <param name="channel">The channel to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public MatchServiceClient(grpc::ChannelBase channel) : base(channel)
      {
      }
      /// <summary>Creates a new client for MatchService that uses a custom <c>CallInvoker</c>.</summary>
      /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public MatchServiceClient(grpc::CallInvoker callInvoker) : base(callInvoker)
      {
      }
      /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected MatchServiceClient() : base()
      {
      }
      /// <summary>Protected constructor to allow creation of configured clients.</summary>
      /// <param name="configuration">The client configuration.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected MatchServiceClient(ClientBaseConfiguration configuration) : base(configuration)
      {
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncServerStreamingCall<global::Thetan.Shared.V1.MatchProtoPackage> CreateMatchWithVersion(global::Thetan.Shared.V1.MatchProtoVersionPackage request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CreateMatchWithVersion(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncServerStreamingCall<global::Thetan.Shared.V1.MatchProtoPackage> CreateMatchWithVersion(global::Thetan.Shared.V1.MatchProtoVersionPackage request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncServerStreamingCall(__Method_CreateMatchWithVersion, null, options, request);
      }
      /// <summary>
      ///register de nhan su kien match found
      /// </summary>
      /// <param name="request">The request to send to the server.</param>
      /// <param name="headers">The initial metadata to send with the call. This parameter is optional.</param>
      /// <param name="deadline">An optional deadline for the call. The call will be cancelled if deadline is hit.</param>
      /// <param name="cancellationToken">An optional token for canceling the call.</param>
      /// <returns>The call object.</returns>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncServerStreamingCall<global::Thetan.Shared.V1.MatchProtoPackage> RegisterMatchFound(global::Thetan.Shared.V1.MatchProtoPackage request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return RegisterMatchFound(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      /// <summary>
      ///register de nhan su kien match found
      /// </summary>
      /// <param name="request">The request to send to the server.</param>
      /// <param name="options">The options for the call.</param>
      /// <returns>The call object.</returns>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncServerStreamingCall<global::Thetan.Shared.V1.MatchProtoPackage> RegisterMatchFound(global::Thetan.Shared.V1.MatchProtoPackage request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncServerStreamingCall(__Method_RegisterMatchFound, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Shared.V1.MatchProtoPackage CancelMatchMaking(global::Thetan.Shared.V1.MatchProtoPackage request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CancelMatchMaking(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Shared.V1.MatchProtoPackage CancelMatchMaking(global::Thetan.Shared.V1.MatchProtoPackage request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_CancelMatchMaking, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Shared.V1.MatchProtoPackage> CancelMatchMakingAsync(global::Thetan.Shared.V1.MatchProtoPackage request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CancelMatchMakingAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Shared.V1.MatchProtoPackage> CancelMatchMakingAsync(global::Thetan.Shared.V1.MatchProtoPackage request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_CancelMatchMaking, null, options, request);
      }
      /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected override MatchServiceClient NewInstance(ClientBaseConfiguration configuration)
      {
        return new MatchServiceClient(configuration);
      }
    }

    /// <summary>Creates service definition that can be registered with a server</summary>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static grpc::ServerServiceDefinition BindService(MatchServiceBase serviceImpl)
    {
      return grpc::ServerServiceDefinition.CreateBuilder()
          .AddMethod(__Method_CreateMatchWithVersion, serviceImpl.CreateMatchWithVersion)
          .AddMethod(__Method_RegisterMatchFound, serviceImpl.RegisterMatchFound)
          .AddMethod(__Method_CancelMatchMaking, serviceImpl.CancelMatchMaking).Build();
    }

    /// <summary>Register service method with a service binder with or without implementation. Useful when customizing the service binding logic.
    /// Note: this method is part of an experimental API that can change or be removed without any prior notice.</summary>
    /// <param name="serviceBinder">Service methods will be bound by calling <c>AddMethod</c> on this object.</param>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static void BindService(grpc::ServiceBinderBase serviceBinder, MatchServiceBase serviceImpl)
    {
      serviceBinder.AddMethod(__Method_CreateMatchWithVersion, serviceImpl == null ? null : new grpc::ServerStreamingServerMethod<global::Thetan.Shared.V1.MatchProtoVersionPackage, global::Thetan.Shared.V1.MatchProtoPackage>(serviceImpl.CreateMatchWithVersion));
      serviceBinder.AddMethod(__Method_RegisterMatchFound, serviceImpl == null ? null : new grpc::ServerStreamingServerMethod<global::Thetan.Shared.V1.MatchProtoPackage, global::Thetan.Shared.V1.MatchProtoPackage>(serviceImpl.RegisterMatchFound));
      serviceBinder.AddMethod(__Method_CancelMatchMaking, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Thetan.Shared.V1.MatchProtoPackage, global::Thetan.Shared.V1.MatchProtoPackage>(serviceImpl.CancelMatchMaking));
    }

  }
  public static partial class MatchHandleService
  {
    static readonly string __ServiceName = "thetan.shared.v1.MatchHandleService";

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
    static readonly grpc::Marshaller<global::Thetan.Shared.V1.DataPlayAgainSuccess> __Marshaller_thetan_shared_v1_DataPlayAgainSuccess = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Shared.V1.DataPlayAgainSuccess.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.Shared.V1.EmptyResponse> __Marshaller_thetan_shared_v1_EmptyResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Shared.V1.EmptyResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.Shared.V1.DeleteTicketSuccess> __Marshaller_thetan_shared_v1_DeleteTicketSuccess = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Shared.V1.DeleteTicketSuccess.Parser));

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Thetan.Shared.V1.DataPlayAgainSuccess, global::Thetan.Shared.V1.EmptyResponse> __Method_HandlePlayAgain = new grpc::Method<global::Thetan.Shared.V1.DataPlayAgainSuccess, global::Thetan.Shared.V1.EmptyResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "HandlePlayAgain",
        __Marshaller_thetan_shared_v1_DataPlayAgainSuccess,
        __Marshaller_thetan_shared_v1_EmptyResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Thetan.Shared.V1.DeleteTicketSuccess, global::Thetan.Shared.V1.EmptyResponse> __Method_HandleDeleteTicket = new grpc::Method<global::Thetan.Shared.V1.DeleteTicketSuccess, global::Thetan.Shared.V1.EmptyResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "HandleDeleteTicket",
        __Marshaller_thetan_shared_v1_DeleteTicketSuccess,
        __Marshaller_thetan_shared_v1_EmptyResponse);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
      get { return global::Thetan.Shared.V1.MatchReflection.Descriptor.Services[1]; }
    }

    /// <summary>Base class for server-side implementations of MatchHandleService</summary>
    [grpc::BindServiceMethod(typeof(MatchHandleService), "BindService")]
    public abstract partial class MatchHandleServiceBase
    {
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Thetan.Shared.V1.EmptyResponse> HandlePlayAgain(global::Thetan.Shared.V1.DataPlayAgainSuccess request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Thetan.Shared.V1.EmptyResponse> HandleDeleteTicket(global::Thetan.Shared.V1.DeleteTicketSuccess request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

    }

    /// <summary>Client for MatchHandleService</summary>
    public partial class MatchHandleServiceClient : grpc::ClientBase<MatchHandleServiceClient>
    {
      /// <summary>Creates a new client for MatchHandleService</summary>
      /// <param name="channel">The channel to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public MatchHandleServiceClient(grpc::ChannelBase channel) : base(channel)
      {
      }
      /// <summary>Creates a new client for MatchHandleService that uses a custom <c>CallInvoker</c>.</summary>
      /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public MatchHandleServiceClient(grpc::CallInvoker callInvoker) : base(callInvoker)
      {
      }
      /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected MatchHandleServiceClient() : base()
      {
      }
      /// <summary>Protected constructor to allow creation of configured clients.</summary>
      /// <param name="configuration">The client configuration.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected MatchHandleServiceClient(ClientBaseConfiguration configuration) : base(configuration)
      {
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Shared.V1.EmptyResponse HandlePlayAgain(global::Thetan.Shared.V1.DataPlayAgainSuccess request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return HandlePlayAgain(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Shared.V1.EmptyResponse HandlePlayAgain(global::Thetan.Shared.V1.DataPlayAgainSuccess request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_HandlePlayAgain, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Shared.V1.EmptyResponse> HandlePlayAgainAsync(global::Thetan.Shared.V1.DataPlayAgainSuccess request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return HandlePlayAgainAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Shared.V1.EmptyResponse> HandlePlayAgainAsync(global::Thetan.Shared.V1.DataPlayAgainSuccess request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_HandlePlayAgain, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Shared.V1.EmptyResponse HandleDeleteTicket(global::Thetan.Shared.V1.DeleteTicketSuccess request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return HandleDeleteTicket(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Shared.V1.EmptyResponse HandleDeleteTicket(global::Thetan.Shared.V1.DeleteTicketSuccess request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_HandleDeleteTicket, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Shared.V1.EmptyResponse> HandleDeleteTicketAsync(global::Thetan.Shared.V1.DeleteTicketSuccess request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return HandleDeleteTicketAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Shared.V1.EmptyResponse> HandleDeleteTicketAsync(global::Thetan.Shared.V1.DeleteTicketSuccess request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_HandleDeleteTicket, null, options, request);
      }
      /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected override MatchHandleServiceClient NewInstance(ClientBaseConfiguration configuration)
      {
        return new MatchHandleServiceClient(configuration);
      }
    }

    /// <summary>Creates service definition that can be registered with a server</summary>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static grpc::ServerServiceDefinition BindService(MatchHandleServiceBase serviceImpl)
    {
      return grpc::ServerServiceDefinition.CreateBuilder()
          .AddMethod(__Method_HandlePlayAgain, serviceImpl.HandlePlayAgain)
          .AddMethod(__Method_HandleDeleteTicket, serviceImpl.HandleDeleteTicket).Build();
    }

    /// <summary>Register service method with a service binder with or without implementation. Useful when customizing the service binding logic.
    /// Note: this method is part of an experimental API that can change or be removed without any prior notice.</summary>
    /// <param name="serviceBinder">Service methods will be bound by calling <c>AddMethod</c> on this object.</param>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static void BindService(grpc::ServiceBinderBase serviceBinder, MatchHandleServiceBase serviceImpl)
    {
      serviceBinder.AddMethod(__Method_HandlePlayAgain, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Thetan.Shared.V1.DataPlayAgainSuccess, global::Thetan.Shared.V1.EmptyResponse>(serviceImpl.HandlePlayAgain));
      serviceBinder.AddMethod(__Method_HandleDeleteTicket, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Thetan.Shared.V1.DeleteTicketSuccess, global::Thetan.Shared.V1.EmptyResponse>(serviceImpl.HandleDeleteTicket));
    }

  }
}
#endregion
