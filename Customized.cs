// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: thetan/shared/v1/customized.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Thetan.Shared.V1 {

  /// <summary>Holder for reflection information generated from thetan/shared/v1/customized.proto</summary>
  public static partial class CustomizedReflection {

    #region Descriptor
    /// <summary>File descriptor for thetan/shared/v1/customized.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static CustomizedReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CiF0aGV0YW4vc2hhcmVkL3YxL2N1c3RvbWl6ZWQucHJvdG8SEHRoZXRhbi5z",
            "aGFyZWQudjEaGWdvb2dsZS9wcm90b2J1Zi9hbnkucHJvdG8aE3RhZ2dlci90",
            "YWdnZXIucHJvdG8ivQIKCkN1c3RvbWl6ZWQSJAoEbmFtZRgBIAEoCUIQmoSe",
            "Awtic29uOiJuYW1lIlIEbmFtZRI2Cgp0ZXh0dXJlVVJMGAIgASgJQhaahJ4D",
            "EWJzb246InRleHR1cmVVUkwiUgp0ZXh0dXJlVVJMEjwKDHRodW1ibmFpbFVS",
            "TBgFIAEoCUIYmoSeAxNic29uOiJ0aHVtYm5haWxVUkwiUgx0aHVtYm5haWxV",
            "UkwSMwoJcHVibGlzaEF0GAMgASgDQhWahJ4DEGJzb246InB1Ymxpc2hBdCJS",
            "CXB1Ymxpc2hBdBJTCgZzaGFkZXIYBCABKAsyIi50aGV0YW4uc2hhcmVkLnYx",
            "LkN1c3RvbWl6ZWRTaGFkZXJCEpqEngMNYnNvbjoic2hhZGVyIkgAUgZzaGFk",
            "ZXKIAQFCCQoHX3NoYWRlciLxAQoQQ3VzdG9taXplZFNoYWRlchIwCghzaGFk",
            "ZXJJRBgBIAEoBUIUmoSeAw9ic29uOiJzaGFkZXJJRCJSCHNoYWRlcklEEloK",
            "BnBhcmFtcxgCIAMoCzIuLnRoZXRhbi5zaGFyZWQudjEuQ3VzdG9taXplZFNo",
            "YWRlci5QYXJhbXNFbnRyeUISmoSeAw1ic29uOiJwYXJhbXMiUgZwYXJhbXMa",
            "TwoLUGFyYW1zRW50cnkSEAoDa2V5GAEgASgJUgNrZXkSKgoFdmFsdWUYAiAB",
            "KAsyFC5nb29nbGUucHJvdG9idWYuQW55UgV2YWx1ZToCOAEiyAEKBEl0ZW0S",
            "JAoEdHlwZRgBIAEoBUIQmoSeAwtic29uOiJ0eXBlIlIEdHlwZRIkCgRraW5k",
            "GAIgASgDQhCahJ4DC2Jzb246ImtpbmQiUgRraW5kEh4KAmlkGAMgASgJQg6a",
            "hJ4DCWJzb246ImlkIlICaWQSVAoKY3VzdG9taXplZBgEIAEoCzIcLnRoZXRh",
            "bi5zaGFyZWQudjEuQ3VzdG9taXplZEIWmoSeAxFic29uOiJjdXN0b21pemVk",
            "IlIKY3VzdG9taXplZEK+AQoUY29tLnRoZXRhbi5zaGFyZWQudjFCD0N1c3Rv",
            "bWl6ZWRQcm90b1ABWjN0aGV0YW4tYnVmL2dlbi9nby90aGV0YW4vc2hhcmVk",
            "L3YxO3RoZXRhbl9zaGFyZWRfdjGiAgNUU1iqAhBUaGV0YW4uU2hhcmVkLlYx",
            "ygIQVGhldGFuXFNoYXJlZFxWMeICHFRoZXRhblxTaGFyZWRcVjFcR1BCTWV0",
            "YWRhdGHqAhJUaGV0YW46OlNoYXJlZDo6VjFiBnByb3RvMw=="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Google.Protobuf.WellKnownTypes.AnyReflection.Descriptor, global::Tagger.TaggerReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Thetan.Shared.V1.Customized), global::Thetan.Shared.V1.Customized.Parser, new[]{ "Name", "TextureURL", "ThumbnailURL", "PublishAt", "Shader" }, new[]{ "Shader" }, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Thetan.Shared.V1.CustomizedShader), global::Thetan.Shared.V1.CustomizedShader.Parser, new[]{ "ShaderID", "Params" }, null, null, null, new pbr::GeneratedClrTypeInfo[] { null, }),
            new pbr::GeneratedClrTypeInfo(typeof(global::Thetan.Shared.V1.Item), global::Thetan.Shared.V1.Item.Parser, new[]{ "Type", "Kind", "Id", "Customized" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  /// <summary>
  /// Customized dùng để gắn vào các item để cho phép app Thetan Creator có thể lưu được bản vẽ.
  /// </summary>
  public sealed partial class Customized : pb::IMessage<Customized>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<Customized> _parser = new pb::MessageParser<Customized>(() => new Customized());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<Customized> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Thetan.Shared.V1.CustomizedReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Customized() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Customized(Customized other) : this() {
      name_ = other.name_;
      textureURL_ = other.textureURL_;
      thumbnailURL_ = other.thumbnailURL_;
      publishAt_ = other.publishAt_;
      shader_ = other.shader_ != null ? other.shader_.Clone() : null;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Customized Clone() {
      return new Customized(this);
    }

    /// <summary>Field number for the "name" field.</summary>
    public const int NameFieldNumber = 1;
    private string name_ = "";
    /// <summary>
    /// Name cho phép user tự đặt tên cho item của mình ở Thetan Creator
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public string Name {
      get { return name_; }
      set {
        name_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "textureURL" field.</summary>
    public const int TextureURLFieldNumber = 2;
    private string textureURL_ = "";
    /// <summary>
    /// TextureURL đường dẫn tuyệt đối đến file texture lưu trên cloud
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public string TextureURL {
      get { return textureURL_; }
      set {
        textureURL_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "thumbnailURL" field.</summary>
    public const int ThumbnailURLFieldNumber = 5;
    private string thumbnailURL_ = "";
    /// <summary>
    /// ThumbnailURL đường dẫn tuyệt đối đến file thumbnail lưu trên cloud
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public string ThumbnailURL {
      get { return thumbnailURL_; }
      set {
        thumbnailURL_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "publishAt" field.</summary>
    public const int PublishAtFieldNumber = 3;
    private long publishAt_;
    /// <summary>
    /// PublishAt thời điểm publish item customized này
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public long PublishAt {
      get { return publishAt_; }
      set {
        publishAt_ = value;
      }
    }

    /// <summary>Field number for the "shader" field.</summary>
    public const int ShaderFieldNumber = 4;
    private global::Thetan.Shared.V1.CustomizedShader shader_;
    /// <summary>
    /// CustomShader shader và các cấu hình của shader đó
    /// </summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public global::Thetan.Shared.V1.CustomizedShader Shader {
      get { return shader_; }
      set {
        shader_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as Customized);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(Customized other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Name != other.Name) return false;
      if (TextureURL != other.TextureURL) return false;
      if (ThumbnailURL != other.ThumbnailURL) return false;
      if (PublishAt != other.PublishAt) return false;
      if (!object.Equals(Shader, other.Shader)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (Name.Length != 0) hash ^= Name.GetHashCode();
      if (TextureURL.Length != 0) hash ^= TextureURL.GetHashCode();
      if (ThumbnailURL.Length != 0) hash ^= ThumbnailURL.GetHashCode();
      if (PublishAt != 0L) hash ^= PublishAt.GetHashCode();
      if (shader_ != null) hash ^= Shader.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      if (Name.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(Name);
      }
      if (TextureURL.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(TextureURL);
      }
      if (PublishAt != 0L) {
        output.WriteRawTag(24);
        output.WriteInt64(PublishAt);
      }
      if (shader_ != null) {
        output.WriteRawTag(34);
        output.WriteMessage(Shader);
      }
      if (ThumbnailURL.Length != 0) {
        output.WriteRawTag(42);
        output.WriteString(ThumbnailURL);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (Name.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(Name);
      }
      if (TextureURL.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(TextureURL);
      }
      if (PublishAt != 0L) {
        output.WriteRawTag(24);
        output.WriteInt64(PublishAt);
      }
      if (shader_ != null) {
        output.WriteRawTag(34);
        output.WriteMessage(Shader);
      }
      if (ThumbnailURL.Length != 0) {
        output.WriteRawTag(42);
        output.WriteString(ThumbnailURL);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int CalculateSize() {
      int size = 0;
      if (Name.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Name);
      }
      if (TextureURL.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(TextureURL);
      }
      if (ThumbnailURL.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(ThumbnailURL);
      }
      if (PublishAt != 0L) {
        size += 1 + pb::CodedOutputStream.ComputeInt64Size(PublishAt);
      }
      if (shader_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Shader);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(Customized other) {
      if (other == null) {
        return;
      }
      if (other.Name.Length != 0) {
        Name = other.Name;
      }
      if (other.TextureURL.Length != 0) {
        TextureURL = other.TextureURL;
      }
      if (other.ThumbnailURL.Length != 0) {
        ThumbnailURL = other.ThumbnailURL;
      }
      if (other.PublishAt != 0L) {
        PublishAt = other.PublishAt;
      }
      if (other.shader_ != null) {
        if (shader_ == null) {
          Shader = new global::Thetan.Shared.V1.CustomizedShader();
        }
        Shader.MergeFrom(other.Shader);
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            Name = input.ReadString();
            break;
          }
          case 18: {
            TextureURL = input.ReadString();
            break;
          }
          case 24: {
            PublishAt = input.ReadInt64();
            break;
          }
          case 34: {
            if (shader_ == null) {
              Shader = new global::Thetan.Shared.V1.CustomizedShader();
            }
            input.ReadMessage(Shader);
            break;
          }
          case 42: {
            ThumbnailURL = input.ReadString();
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 10: {
            Name = input.ReadString();
            break;
          }
          case 18: {
            TextureURL = input.ReadString();
            break;
          }
          case 24: {
            PublishAt = input.ReadInt64();
            break;
          }
          case 34: {
            if (shader_ == null) {
              Shader = new global::Thetan.Shared.V1.CustomizedShader();
            }
            input.ReadMessage(Shader);
            break;
          }
          case 42: {
            ThumbnailURL = input.ReadString();
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class CustomizedShader : pb::IMessage<CustomizedShader>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<CustomizedShader> _parser = new pb::MessageParser<CustomizedShader>(() => new CustomizedShader());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<CustomizedShader> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Thetan.Shared.V1.CustomizedReflection.Descriptor.MessageTypes[1]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public CustomizedShader() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public CustomizedShader(CustomizedShader other) : this() {
      shaderID_ = other.shaderID_;
      params_ = other.params_.Clone();
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public CustomizedShader Clone() {
      return new CustomizedShader(this);
    }

    /// <summary>Field number for the "shaderID" field.</summary>
    public const int ShaderIDFieldNumber = 1;
    private int shaderID_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int ShaderID {
      get { return shaderID_; }
      set {
        shaderID_ = value;
      }
    }

    /// <summary>Field number for the "params" field.</summary>
    public const int ParamsFieldNumber = 2;
    private static readonly pbc::MapField<string, global::Google.Protobuf.WellKnownTypes.Any>.Codec _map_params_codec
        = new pbc::MapField<string, global::Google.Protobuf.WellKnownTypes.Any>.Codec(pb::FieldCodec.ForString(10, ""), pb::FieldCodec.ForMessage(18, global::Google.Protobuf.WellKnownTypes.Any.Parser), 18);
    private readonly pbc::MapField<string, global::Google.Protobuf.WellKnownTypes.Any> params_ = new pbc::MapField<string, global::Google.Protobuf.WellKnownTypes.Any>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public pbc::MapField<string, global::Google.Protobuf.WellKnownTypes.Any> Params {
      get { return params_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as CustomizedShader);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(CustomizedShader other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (ShaderID != other.ShaderID) return false;
      if (!Params.Equals(other.Params)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (ShaderID != 0) hash ^= ShaderID.GetHashCode();
      hash ^= Params.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      if (ShaderID != 0) {
        output.WriteRawTag(8);
        output.WriteInt32(ShaderID);
      }
      params_.WriteTo(output, _map_params_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (ShaderID != 0) {
        output.WriteRawTag(8);
        output.WriteInt32(ShaderID);
      }
      params_.WriteTo(ref output, _map_params_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int CalculateSize() {
      int size = 0;
      if (ShaderID != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(ShaderID);
      }
      size += params_.CalculateSize(_map_params_codec);
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(CustomizedShader other) {
      if (other == null) {
        return;
      }
      if (other.ShaderID != 0) {
        ShaderID = other.ShaderID;
      }
      params_.Add(other.params_);
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 8: {
            ShaderID = input.ReadInt32();
            break;
          }
          case 18: {
            params_.AddEntriesFrom(input, _map_params_codec);
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 8: {
            ShaderID = input.ReadInt32();
            break;
          }
          case 18: {
            params_.AddEntriesFrom(ref input, _map_params_codec);
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class Item : pb::IMessage<Item>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<Item> _parser = new pb::MessageParser<Item>(() => new Item());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<Item> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Thetan.Shared.V1.CustomizedReflection.Descriptor.MessageTypes[2]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Item() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Item(Item other) : this() {
      type_ = other.type_;
      kind_ = other.kind_;
      id_ = other.id_;
      customized_ = other.customized_ != null ? other.customized_.Clone() : null;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Item Clone() {
      return new Item(this);
    }

    /// <summary>Field number for the "type" field.</summary>
    public const int TypeFieldNumber = 1;
    private int type_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int Type {
      get { return type_; }
      set {
        type_ = value;
      }
    }

    /// <summary>Field number for the "kind" field.</summary>
    public const int KindFieldNumber = 2;
    private long kind_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public long Kind {
      get { return kind_; }
      set {
        kind_ = value;
      }
    }

    /// <summary>Field number for the "id" field.</summary>
    public const int IdFieldNumber = 3;
    private string id_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public string Id {
      get { return id_; }
      set {
        id_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "customized" field.</summary>
    public const int CustomizedFieldNumber = 4;
    private global::Thetan.Shared.V1.Customized customized_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public global::Thetan.Shared.V1.Customized Customized {
      get { return customized_; }
      set {
        customized_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as Item);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(Item other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Type != other.Type) return false;
      if (Kind != other.Kind) return false;
      if (Id != other.Id) return false;
      if (!object.Equals(Customized, other.Customized)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (Type != 0) hash ^= Type.GetHashCode();
      if (Kind != 0L) hash ^= Kind.GetHashCode();
      if (Id.Length != 0) hash ^= Id.GetHashCode();
      if (customized_ != null) hash ^= Customized.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      if (Type != 0) {
        output.WriteRawTag(8);
        output.WriteInt32(Type);
      }
      if (Kind != 0L) {
        output.WriteRawTag(16);
        output.WriteInt64(Kind);
      }
      if (Id.Length != 0) {
        output.WriteRawTag(26);
        output.WriteString(Id);
      }
      if (customized_ != null) {
        output.WriteRawTag(34);
        output.WriteMessage(Customized);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (Type != 0) {
        output.WriteRawTag(8);
        output.WriteInt32(Type);
      }
      if (Kind != 0L) {
        output.WriteRawTag(16);
        output.WriteInt64(Kind);
      }
      if (Id.Length != 0) {
        output.WriteRawTag(26);
        output.WriteString(Id);
      }
      if (customized_ != null) {
        output.WriteRawTag(34);
        output.WriteMessage(Customized);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int CalculateSize() {
      int size = 0;
      if (Type != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(Type);
      }
      if (Kind != 0L) {
        size += 1 + pb::CodedOutputStream.ComputeInt64Size(Kind);
      }
      if (Id.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Id);
      }
      if (customized_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Customized);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(Item other) {
      if (other == null) {
        return;
      }
      if (other.Type != 0) {
        Type = other.Type;
      }
      if (other.Kind != 0L) {
        Kind = other.Kind;
      }
      if (other.Id.Length != 0) {
        Id = other.Id;
      }
      if (other.customized_ != null) {
        if (customized_ == null) {
          Customized = new global::Thetan.Shared.V1.Customized();
        }
        Customized.MergeFrom(other.Customized);
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 8: {
            Type = input.ReadInt32();
            break;
          }
          case 16: {
            Kind = input.ReadInt64();
            break;
          }
          case 26: {
            Id = input.ReadString();
            break;
          }
          case 34: {
            if (customized_ == null) {
              Customized = new global::Thetan.Shared.V1.Customized();
            }
            input.ReadMessage(Customized);
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 8: {
            Type = input.ReadInt32();
            break;
          }
          case 16: {
            Kind = input.ReadInt64();
            break;
          }
          case 26: {
            Id = input.ReadString();
            break;
          }
          case 34: {
            if (customized_ == null) {
              Customized = new global::Thetan.Shared.V1.Customized();
            }
            input.ReadMessage(Customized);
            break;
          }
        }
      }
    }
    #endif

  }

  #endregion

}

#endregion Designer generated code
