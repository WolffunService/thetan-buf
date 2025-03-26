// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: thetan/immortal/v1/immortal_game_info.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Thetan.Immortal.V1 {

  /// <summary>Holder for reflection information generated from thetan/immortal/v1/immortal_game_info.proto</summary>
  public static partial class ImmortalGameInfoReflection {

    #region Descriptor
    /// <summary>File descriptor for thetan/immortal/v1/immortal_game_info.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static ImmortalGameInfoReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "Cit0aGV0YW4vaW1tb3J0YWwvdjEvaW1tb3J0YWxfZ2FtZV9pbmZvLnByb3Rv",
            "EhJ0aGV0YW4uaW1tb3J0YWwudjEiUgoQU2V0SW5HYW1lTW9kZU1zZxI+Cgpp",
            "bkdhbWVNb2RlGAEgASgOMh4udGhldGFuLmltbW9ydGFsLnYxLkluR2FtZU1v",
            "ZGVSCmluR2FtZU1vZGUicwoVT3RoZXJTZXRJbkdhbWVNb2RlTXNnEhoKCHBs",
            "YXllcklEGAEgASgJUghwbGF5ZXJJRBI+CgppbkdhbWVNb2RlGAIgASgOMh4u",
            "dGhldGFuLmltbW9ydGFsLnYxLkluR2FtZU1vZGVSCmluR2FtZU1vZGUqQwoI",
            "R2FtZU1vZGUSDQoJTk9ORV9NT0RFEAASCgoGUkFOS0VEEAESDgoKT05CT0FS",
            "RElORxACEgwKCEZSSUVORExZEAMqKAoKSW5HYW1lTW9kZRIICgRTT0xPEAAS",
            "BgoCS08QARIICgREVUFMEAJC0gEKFmNvbS50aGV0YW4uaW1tb3J0YWwudjFC",
            "FUltbW9ydGFsR2FtZUluZm9Qcm90b1ABWjd0aGV0YW4tYnVmL2dlbi9nby90",
            "aGV0YW4vaW1tb3J0YWwvdjE7dGhldGFuX2ltbW9ydGFsX3YxogIDVElYqgIS",
            "VGhldGFuLkltbW9ydGFsLlYxygISVGhldGFuXEltbW9ydGFsXFYx4gIeVGhl",
            "dGFuXEltbW9ydGFsXFYxXEdQQk1ldGFkYXRh6gIUVGhldGFuOjpJbW1vcnRh",
            "bDo6VjFiBnByb3RvMw=="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(new[] {typeof(global::Thetan.Immortal.V1.GameMode), typeof(global::Thetan.Immortal.V1.InGameMode), }, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Thetan.Immortal.V1.SetInGameModeMsg), global::Thetan.Immortal.V1.SetInGameModeMsg.Parser, new[]{ "InGameMode" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Thetan.Immortal.V1.OtherSetInGameModeMsg), global::Thetan.Immortal.V1.OtherSetInGameModeMsg.Parser, new[]{ "PlayerID", "InGameMode" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Enums
  public enum GameMode {
    [pbr::OriginalName("NONE_MODE")] NoneMode = 0,
    [pbr::OriginalName("RANKED")] Ranked = 1,
    [pbr::OriginalName("ONBOARDING")] Onboarding = 2,
    [pbr::OriginalName("FRIENDLY")] Friendly = 3,
  }

  public enum InGameMode {
    [pbr::OriginalName("SOLO")] Solo = 0,
    [pbr::OriginalName("KO")] Ko = 1,
    [pbr::OriginalName("DUAL")] Dual = 2,
  }

  #endregion

  #region Messages
  public sealed partial class SetInGameModeMsg : pb::IMessage<SetInGameModeMsg>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<SetInGameModeMsg> _parser = new pb::MessageParser<SetInGameModeMsg>(() => new SetInGameModeMsg());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<SetInGameModeMsg> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Thetan.Immortal.V1.ImmortalGameInfoReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public SetInGameModeMsg() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public SetInGameModeMsg(SetInGameModeMsg other) : this() {
      inGameMode_ = other.inGameMode_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public SetInGameModeMsg Clone() {
      return new SetInGameModeMsg(this);
    }

    /// <summary>Field number for the "inGameMode" field.</summary>
    public const int InGameModeFieldNumber = 1;
    private global::Thetan.Immortal.V1.InGameMode inGameMode_ = global::Thetan.Immortal.V1.InGameMode.Solo;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public global::Thetan.Immortal.V1.InGameMode InGameMode {
      get { return inGameMode_; }
      set {
        inGameMode_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as SetInGameModeMsg);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(SetInGameModeMsg other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (InGameMode != other.InGameMode) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (InGameMode != global::Thetan.Immortal.V1.InGameMode.Solo) hash ^= InGameMode.GetHashCode();
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
      if (InGameMode != global::Thetan.Immortal.V1.InGameMode.Solo) {
        output.WriteRawTag(8);
        output.WriteEnum((int) InGameMode);
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
      if (InGameMode != global::Thetan.Immortal.V1.InGameMode.Solo) {
        output.WriteRawTag(8);
        output.WriteEnum((int) InGameMode);
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
      if (InGameMode != global::Thetan.Immortal.V1.InGameMode.Solo) {
        size += 1 + pb::CodedOutputStream.ComputeEnumSize((int) InGameMode);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(SetInGameModeMsg other) {
      if (other == null) {
        return;
      }
      if (other.InGameMode != global::Thetan.Immortal.V1.InGameMode.Solo) {
        InGameMode = other.InGameMode;
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
            InGameMode = (global::Thetan.Immortal.V1.InGameMode) input.ReadEnum();
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
            InGameMode = (global::Thetan.Immortal.V1.InGameMode) input.ReadEnum();
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class OtherSetInGameModeMsg : pb::IMessage<OtherSetInGameModeMsg>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<OtherSetInGameModeMsg> _parser = new pb::MessageParser<OtherSetInGameModeMsg>(() => new OtherSetInGameModeMsg());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<OtherSetInGameModeMsg> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Thetan.Immortal.V1.ImmortalGameInfoReflection.Descriptor.MessageTypes[1]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public OtherSetInGameModeMsg() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public OtherSetInGameModeMsg(OtherSetInGameModeMsg other) : this() {
      playerID_ = other.playerID_;
      inGameMode_ = other.inGameMode_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public OtherSetInGameModeMsg Clone() {
      return new OtherSetInGameModeMsg(this);
    }

    /// <summary>Field number for the "playerID" field.</summary>
    public const int PlayerIDFieldNumber = 1;
    private string playerID_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public string PlayerID {
      get { return playerID_; }
      set {
        playerID_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "inGameMode" field.</summary>
    public const int InGameModeFieldNumber = 2;
    private global::Thetan.Immortal.V1.InGameMode inGameMode_ = global::Thetan.Immortal.V1.InGameMode.Solo;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public global::Thetan.Immortal.V1.InGameMode InGameMode {
      get { return inGameMode_; }
      set {
        inGameMode_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as OtherSetInGameModeMsg);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(OtherSetInGameModeMsg other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (PlayerID != other.PlayerID) return false;
      if (InGameMode != other.InGameMode) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (PlayerID.Length != 0) hash ^= PlayerID.GetHashCode();
      if (InGameMode != global::Thetan.Immortal.V1.InGameMode.Solo) hash ^= InGameMode.GetHashCode();
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
      if (PlayerID.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(PlayerID);
      }
      if (InGameMode != global::Thetan.Immortal.V1.InGameMode.Solo) {
        output.WriteRawTag(16);
        output.WriteEnum((int) InGameMode);
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
      if (PlayerID.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(PlayerID);
      }
      if (InGameMode != global::Thetan.Immortal.V1.InGameMode.Solo) {
        output.WriteRawTag(16);
        output.WriteEnum((int) InGameMode);
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
      if (PlayerID.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(PlayerID);
      }
      if (InGameMode != global::Thetan.Immortal.V1.InGameMode.Solo) {
        size += 1 + pb::CodedOutputStream.ComputeEnumSize((int) InGameMode);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(OtherSetInGameModeMsg other) {
      if (other == null) {
        return;
      }
      if (other.PlayerID.Length != 0) {
        PlayerID = other.PlayerID;
      }
      if (other.InGameMode != global::Thetan.Immortal.V1.InGameMode.Solo) {
        InGameMode = other.InGameMode;
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
            PlayerID = input.ReadString();
            break;
          }
          case 16: {
            InGameMode = (global::Thetan.Immortal.V1.InGameMode) input.ReadEnum();
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
            PlayerID = input.ReadString();
            break;
          }
          case 16: {
            InGameMode = (global::Thetan.Immortal.V1.InGameMode) input.ReadEnum();
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
