// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: thetan/multiplayer/v1/storage.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Thetan.Multiplayer.V1 {

  /// <summary>Holder for reflection information generated from thetan/multiplayer/v1/storage.proto</summary>
  public static partial class StorageReflection {

    #region Descriptor
    /// <summary>File descriptor for thetan/multiplayer/v1/storage.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static StorageReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CiN0aGV0YW4vbXVsdGlwbGF5ZXIvdjEvc3RvcmFnZS5wcm90bxIVdGhldGFu",
            "Lm11bHRpcGxheWVyLnYxGiJ0aGV0YW4vbXVsdGlwbGF5ZXIvdjEvY29tbW9u",
            "LnByb3RvIrgBChRQbGF5ZXJTdG9yYWdlTWVzc2FnZRJGCgpwbGF5ZXJJbmZv",
            "GAEgASgLMiYudGhldGFuLm11bHRpcGxheWVyLnYxLlBsYXllckluZm9Qcm90",
            "b1IKcGxheWVySW5mbxIYCgd2ZXJzaW9uGAIgASgFUgd2ZXJzaW9uEj4KCW1h",
            "dGNoSW5mbxgDIAEoCzIgLnRoZXRhbi5tdWx0aXBsYXllci52MS5NYXRjaElu",
            "Zm9SCW1hdGNoSW5mbyLwAwoTUGFydHlTdG9yYWdlTWVzc2FnZRJDCglwYXJ0",
            "eUluZm8YASABKAsyJS50aGV0YW4ubXVsdGlwbGF5ZXIudjEuUGFydHlJbmZv",
            "UHJvdG9SCXBhcnR5SW5mbxJ+ChZwbGF5ZXJzUmVhZHlGb3JSZW1hdGNoGAIg",
            "AygLMkYudGhldGFuLm11bHRpcGxheWVyLnYxLlBhcnR5U3RvcmFnZU1lc3Nh",
            "Z2UuUGxheWVyc1JlYWR5Rm9yUmVtYXRjaEVudHJ5UhZwbGF5ZXJzUmVhZHlG",
            "b3JSZW1hdGNoEmYKDnBsYXllcnNSZWdpb25zGAMgAygLMj4udGhldGFuLm11",
            "bHRpcGxheWVyLnYxLlBhcnR5U3RvcmFnZU1lc3NhZ2UuUGxheWVyc1JlZ2lv",
            "bnNFbnRyeVIOcGxheWVyc1JlZ2lvbnMaSQobUGxheWVyc1JlYWR5Rm9yUmVt",
            "YXRjaEVudHJ5EhAKA2tleRgBIAEoCVIDa2V5EhQKBXZhbHVlGAIgASgIUgV2",
            "YWx1ZToCOAEaYQoTUGxheWVyc1JlZ2lvbnNFbnRyeRIQCgNrZXkYASABKAlS",
            "A2tleRI0CgV2YWx1ZRgCIAEoCzIeLnRoZXRhbi5tdWx0aXBsYXllci52MS5S",
            "ZWdpb25zUgV2YWx1ZToCOAEiIwoHUmVnaW9ucxIYCgdyZWdpb25zGAEgAygF",
            "UgdyZWdpb25zIkEKCU1hdGNoSW5mbxIaCgh0aWNrZXRJRBgBIAEoCVIIdGlj",
            "a2V0SUQSGAoHbWF0Y2hJRBgCIAEoCVIHbWF0Y2hJRELeAQoZY29tLnRoZXRh",
            "bi5tdWx0aXBsYXllci52MUIMU3RvcmFnZVByb3RvUAFaPXRoZXRhbi1idWYv",
            "Z2VuL2dvL3RoZXRhbi9tdWx0aXBsYXllci92MTt0aGV0YW5fbXVsdGlwbGF5",
            "ZXJfdjGiAgNUTViqAhVUaGV0YW4uTXVsdGlwbGF5ZXIuVjHKAhVUaGV0YW5c",
            "TXVsdGlwbGF5ZXJcVjHiAiFUaGV0YW5cTXVsdGlwbGF5ZXJcVjFcR1BCTWV0",
            "YWRhdGHqAhdUaGV0YW46Ok11bHRpcGxheWVyOjpWMWIGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Thetan.Multiplayer.V1.CommonReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Thetan.Multiplayer.V1.PlayerStorageMessage), global::Thetan.Multiplayer.V1.PlayerStorageMessage.Parser, new[]{ "PlayerInfo", "Version", "MatchInfo" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Thetan.Multiplayer.V1.PartyStorageMessage), global::Thetan.Multiplayer.V1.PartyStorageMessage.Parser, new[]{ "PartyInfo", "PlayersReadyForRematch", "PlayersRegions" }, null, null, null, new pbr::GeneratedClrTypeInfo[] { null, null, }),
            new pbr::GeneratedClrTypeInfo(typeof(global::Thetan.Multiplayer.V1.Regions), global::Thetan.Multiplayer.V1.Regions.Parser, new[]{ "Regions_" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Thetan.Multiplayer.V1.MatchInfo), global::Thetan.Multiplayer.V1.MatchInfo.Parser, new[]{ "TicketID", "MatchID" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class PlayerStorageMessage : pb::IMessage<PlayerStorageMessage>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<PlayerStorageMessage> _parser = new pb::MessageParser<PlayerStorageMessage>(() => new PlayerStorageMessage());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<PlayerStorageMessage> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Thetan.Multiplayer.V1.StorageReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public PlayerStorageMessage() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public PlayerStorageMessage(PlayerStorageMessage other) : this() {
      playerInfo_ = other.playerInfo_ != null ? other.playerInfo_.Clone() : null;
      version_ = other.version_;
      matchInfo_ = other.matchInfo_ != null ? other.matchInfo_.Clone() : null;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public PlayerStorageMessage Clone() {
      return new PlayerStorageMessage(this);
    }

    /// <summary>Field number for the "playerInfo" field.</summary>
    public const int PlayerInfoFieldNumber = 1;
    private global::Thetan.Multiplayer.V1.PlayerInfoProto playerInfo_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public global::Thetan.Multiplayer.V1.PlayerInfoProto PlayerInfo {
      get { return playerInfo_; }
      set {
        playerInfo_ = value;
      }
    }

    /// <summary>Field number for the "version" field.</summary>
    public const int VersionFieldNumber = 2;
    private int version_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int Version {
      get { return version_; }
      set {
        version_ = value;
      }
    }

    /// <summary>Field number for the "matchInfo" field.</summary>
    public const int MatchInfoFieldNumber = 3;
    private global::Thetan.Multiplayer.V1.MatchInfo matchInfo_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public global::Thetan.Multiplayer.V1.MatchInfo MatchInfo {
      get { return matchInfo_; }
      set {
        matchInfo_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as PlayerStorageMessage);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(PlayerStorageMessage other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (!object.Equals(PlayerInfo, other.PlayerInfo)) return false;
      if (Version != other.Version) return false;
      if (!object.Equals(MatchInfo, other.MatchInfo)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (playerInfo_ != null) hash ^= PlayerInfo.GetHashCode();
      if (Version != 0) hash ^= Version.GetHashCode();
      if (matchInfo_ != null) hash ^= MatchInfo.GetHashCode();
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
      if (playerInfo_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(PlayerInfo);
      }
      if (Version != 0) {
        output.WriteRawTag(16);
        output.WriteInt32(Version);
      }
      if (matchInfo_ != null) {
        output.WriteRawTag(26);
        output.WriteMessage(MatchInfo);
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
      if (playerInfo_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(PlayerInfo);
      }
      if (Version != 0) {
        output.WriteRawTag(16);
        output.WriteInt32(Version);
      }
      if (matchInfo_ != null) {
        output.WriteRawTag(26);
        output.WriteMessage(MatchInfo);
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
      if (playerInfo_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(PlayerInfo);
      }
      if (Version != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(Version);
      }
      if (matchInfo_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(MatchInfo);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(PlayerStorageMessage other) {
      if (other == null) {
        return;
      }
      if (other.playerInfo_ != null) {
        if (playerInfo_ == null) {
          PlayerInfo = new global::Thetan.Multiplayer.V1.PlayerInfoProto();
        }
        PlayerInfo.MergeFrom(other.PlayerInfo);
      }
      if (other.Version != 0) {
        Version = other.Version;
      }
      if (other.matchInfo_ != null) {
        if (matchInfo_ == null) {
          MatchInfo = new global::Thetan.Multiplayer.V1.MatchInfo();
        }
        MatchInfo.MergeFrom(other.MatchInfo);
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
            if (playerInfo_ == null) {
              PlayerInfo = new global::Thetan.Multiplayer.V1.PlayerInfoProto();
            }
            input.ReadMessage(PlayerInfo);
            break;
          }
          case 16: {
            Version = input.ReadInt32();
            break;
          }
          case 26: {
            if (matchInfo_ == null) {
              MatchInfo = new global::Thetan.Multiplayer.V1.MatchInfo();
            }
            input.ReadMessage(MatchInfo);
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
            if (playerInfo_ == null) {
              PlayerInfo = new global::Thetan.Multiplayer.V1.PlayerInfoProto();
            }
            input.ReadMessage(PlayerInfo);
            break;
          }
          case 16: {
            Version = input.ReadInt32();
            break;
          }
          case 26: {
            if (matchInfo_ == null) {
              MatchInfo = new global::Thetan.Multiplayer.V1.MatchInfo();
            }
            input.ReadMessage(MatchInfo);
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class PartyStorageMessage : pb::IMessage<PartyStorageMessage>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<PartyStorageMessage> _parser = new pb::MessageParser<PartyStorageMessage>(() => new PartyStorageMessage());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<PartyStorageMessage> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Thetan.Multiplayer.V1.StorageReflection.Descriptor.MessageTypes[1]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public PartyStorageMessage() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public PartyStorageMessage(PartyStorageMessage other) : this() {
      partyInfo_ = other.partyInfo_ != null ? other.partyInfo_.Clone() : null;
      playersReadyForRematch_ = other.playersReadyForRematch_.Clone();
      playersRegions_ = other.playersRegions_.Clone();
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public PartyStorageMessage Clone() {
      return new PartyStorageMessage(this);
    }

    /// <summary>Field number for the "partyInfo" field.</summary>
    public const int PartyInfoFieldNumber = 1;
    private global::Thetan.Multiplayer.V1.PartyInfoProto partyInfo_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public global::Thetan.Multiplayer.V1.PartyInfoProto PartyInfo {
      get { return partyInfo_; }
      set {
        partyInfo_ = value;
      }
    }

    /// <summary>Field number for the "playersReadyForRematch" field.</summary>
    public const int PlayersReadyForRematchFieldNumber = 2;
    private static readonly pbc::MapField<string, bool>.Codec _map_playersReadyForRematch_codec
        = new pbc::MapField<string, bool>.Codec(pb::FieldCodec.ForString(10, ""), pb::FieldCodec.ForBool(16, false), 18);
    private readonly pbc::MapField<string, bool> playersReadyForRematch_ = new pbc::MapField<string, bool>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public pbc::MapField<string, bool> PlayersReadyForRematch {
      get { return playersReadyForRematch_; }
    }

    /// <summary>Field number for the "playersRegions" field.</summary>
    public const int PlayersRegionsFieldNumber = 3;
    private static readonly pbc::MapField<string, global::Thetan.Multiplayer.V1.Regions>.Codec _map_playersRegions_codec
        = new pbc::MapField<string, global::Thetan.Multiplayer.V1.Regions>.Codec(pb::FieldCodec.ForString(10, ""), pb::FieldCodec.ForMessage(18, global::Thetan.Multiplayer.V1.Regions.Parser), 26);
    private readonly pbc::MapField<string, global::Thetan.Multiplayer.V1.Regions> playersRegions_ = new pbc::MapField<string, global::Thetan.Multiplayer.V1.Regions>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public pbc::MapField<string, global::Thetan.Multiplayer.V1.Regions> PlayersRegions {
      get { return playersRegions_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as PartyStorageMessage);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(PartyStorageMessage other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (!object.Equals(PartyInfo, other.PartyInfo)) return false;
      if (!PlayersReadyForRematch.Equals(other.PlayersReadyForRematch)) return false;
      if (!PlayersRegions.Equals(other.PlayersRegions)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (partyInfo_ != null) hash ^= PartyInfo.GetHashCode();
      hash ^= PlayersReadyForRematch.GetHashCode();
      hash ^= PlayersRegions.GetHashCode();
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
      if (partyInfo_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(PartyInfo);
      }
      playersReadyForRematch_.WriteTo(output, _map_playersReadyForRematch_codec);
      playersRegions_.WriteTo(output, _map_playersRegions_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (partyInfo_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(PartyInfo);
      }
      playersReadyForRematch_.WriteTo(ref output, _map_playersReadyForRematch_codec);
      playersRegions_.WriteTo(ref output, _map_playersRegions_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int CalculateSize() {
      int size = 0;
      if (partyInfo_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(PartyInfo);
      }
      size += playersReadyForRematch_.CalculateSize(_map_playersReadyForRematch_codec);
      size += playersRegions_.CalculateSize(_map_playersRegions_codec);
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(PartyStorageMessage other) {
      if (other == null) {
        return;
      }
      if (other.partyInfo_ != null) {
        if (partyInfo_ == null) {
          PartyInfo = new global::Thetan.Multiplayer.V1.PartyInfoProto();
        }
        PartyInfo.MergeFrom(other.PartyInfo);
      }
      playersReadyForRematch_.Add(other.playersReadyForRematch_);
      playersRegions_.Add(other.playersRegions_);
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
            if (partyInfo_ == null) {
              PartyInfo = new global::Thetan.Multiplayer.V1.PartyInfoProto();
            }
            input.ReadMessage(PartyInfo);
            break;
          }
          case 18: {
            playersReadyForRematch_.AddEntriesFrom(input, _map_playersReadyForRematch_codec);
            break;
          }
          case 26: {
            playersRegions_.AddEntriesFrom(input, _map_playersRegions_codec);
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
            if (partyInfo_ == null) {
              PartyInfo = new global::Thetan.Multiplayer.V1.PartyInfoProto();
            }
            input.ReadMessage(PartyInfo);
            break;
          }
          case 18: {
            playersReadyForRematch_.AddEntriesFrom(ref input, _map_playersReadyForRematch_codec);
            break;
          }
          case 26: {
            playersRegions_.AddEntriesFrom(ref input, _map_playersRegions_codec);
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class Regions : pb::IMessage<Regions>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<Regions> _parser = new pb::MessageParser<Regions>(() => new Regions());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<Regions> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Thetan.Multiplayer.V1.StorageReflection.Descriptor.MessageTypes[2]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Regions() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Regions(Regions other) : this() {
      regions_ = other.regions_.Clone();
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Regions Clone() {
      return new Regions(this);
    }

    /// <summary>Field number for the "regions" field.</summary>
    public const int Regions_FieldNumber = 1;
    private static readonly pb::FieldCodec<int> _repeated_regions_codec
        = pb::FieldCodec.ForInt32(10);
    private readonly pbc::RepeatedField<int> regions_ = new pbc::RepeatedField<int>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public pbc::RepeatedField<int> Regions_ {
      get { return regions_; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as Regions);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(Regions other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if(!regions_.Equals(other.regions_)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      hash ^= regions_.GetHashCode();
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
      regions_.WriteTo(output, _repeated_regions_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      regions_.WriteTo(ref output, _repeated_regions_codec);
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int CalculateSize() {
      int size = 0;
      size += regions_.CalculateSize(_repeated_regions_codec);
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(Regions other) {
      if (other == null) {
        return;
      }
      regions_.Add(other.regions_);
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
          case 10:
          case 8: {
            regions_.AddEntriesFrom(input, _repeated_regions_codec);
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
          case 10:
          case 8: {
            regions_.AddEntriesFrom(ref input, _repeated_regions_codec);
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class MatchInfo : pb::IMessage<MatchInfo>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<MatchInfo> _parser = new pb::MessageParser<MatchInfo>(() => new MatchInfo());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<MatchInfo> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Thetan.Multiplayer.V1.StorageReflection.Descriptor.MessageTypes[3]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public MatchInfo() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public MatchInfo(MatchInfo other) : this() {
      ticketID_ = other.ticketID_;
      matchID_ = other.matchID_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public MatchInfo Clone() {
      return new MatchInfo(this);
    }

    /// <summary>Field number for the "ticketID" field.</summary>
    public const int TicketIDFieldNumber = 1;
    private string ticketID_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public string TicketID {
      get { return ticketID_; }
      set {
        ticketID_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "matchID" field.</summary>
    public const int MatchIDFieldNumber = 2;
    private string matchID_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public string MatchID {
      get { return matchID_; }
      set {
        matchID_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as MatchInfo);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(MatchInfo other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (TicketID != other.TicketID) return false;
      if (MatchID != other.MatchID) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (TicketID.Length != 0) hash ^= TicketID.GetHashCode();
      if (MatchID.Length != 0) hash ^= MatchID.GetHashCode();
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
      if (TicketID.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(TicketID);
      }
      if (MatchID.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(MatchID);
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
      if (TicketID.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(TicketID);
      }
      if (MatchID.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(MatchID);
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
      if (TicketID.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(TicketID);
      }
      if (MatchID.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(MatchID);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(MatchInfo other) {
      if (other == null) {
        return;
      }
      if (other.TicketID.Length != 0) {
        TicketID = other.TicketID;
      }
      if (other.MatchID.Length != 0) {
        MatchID = other.MatchID;
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
            TicketID = input.ReadString();
            break;
          }
          case 18: {
            MatchID = input.ReadString();
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
            TicketID = input.ReadString();
            break;
          }
          case 18: {
            MatchID = input.ReadString();
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
