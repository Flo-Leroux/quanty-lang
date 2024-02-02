#include <tree_sitter/parser.h>

#if defined(__GNUC__) || defined(__clang__)
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wmissing-field-initializers"
#endif

#define LANGUAGE_VERSION 14
#define STATE_COUNT 16
#define LARGE_STATE_COUNT 2
#define SYMBOL_COUNT 13
#define ALIAS_COUNT 0
#define TOKEN_COUNT 6
#define EXTERNAL_TOKEN_COUNT 3
#define FIELD_COUNT 3
#define MAX_ALIAS_SEQUENCE_LENGTH 5
#define PRODUCTION_ID_COUNT 4

enum {
  anon_sym_component = 1,
  sym_identifier = 2,
  sym__newline = 3,
  sym__indent = 4,
  sym__dedent = 5,
  sym_program = 6,
  sym__fragment = 7,
  sym_component_declaration = 8,
  sym_tag = 9,
  sym_children = 10,
  aux_sym_program_repeat1 = 11,
  aux_sym_component_declaration_repeat1 = 12,
};

static const char * const ts_symbol_names[] = {
  [ts_builtin_sym_end] = "end",
  [anon_sym_component] = "component",
  [sym_identifier] = "identifier",
  [sym__newline] = "_newline",
  [sym__indent] = "_indent",
  [sym__dedent] = "_dedent",
  [sym_program] = "program",
  [sym__fragment] = "_fragment",
  [sym_component_declaration] = "component_declaration",
  [sym_tag] = "tag",
  [sym_children] = "children",
  [aux_sym_program_repeat1] = "program_repeat1",
  [aux_sym_component_declaration_repeat1] = "component_declaration_repeat1",
};

static const TSSymbol ts_symbol_map[] = {
  [ts_builtin_sym_end] = ts_builtin_sym_end,
  [anon_sym_component] = anon_sym_component,
  [sym_identifier] = sym_identifier,
  [sym__newline] = sym__newline,
  [sym__indent] = sym__indent,
  [sym__dedent] = sym__dedent,
  [sym_program] = sym_program,
  [sym__fragment] = sym__fragment,
  [sym_component_declaration] = sym_component_declaration,
  [sym_tag] = sym_tag,
  [sym_children] = sym_children,
  [aux_sym_program_repeat1] = aux_sym_program_repeat1,
  [aux_sym_component_declaration_repeat1] = aux_sym_component_declaration_repeat1,
};

static const TSSymbolMetadata ts_symbol_metadata[] = {
  [ts_builtin_sym_end] = {
    .visible = false,
    .named = true,
  },
  [anon_sym_component] = {
    .visible = true,
    .named = false,
  },
  [sym_identifier] = {
    .visible = true,
    .named = true,
  },
  [sym__newline] = {
    .visible = false,
    .named = true,
  },
  [sym__indent] = {
    .visible = false,
    .named = true,
  },
  [sym__dedent] = {
    .visible = false,
    .named = true,
  },
  [sym_program] = {
    .visible = true,
    .named = true,
  },
  [sym__fragment] = {
    .visible = false,
    .named = true,
  },
  [sym_component_declaration] = {
    .visible = true,
    .named = true,
  },
  [sym_tag] = {
    .visible = true,
    .named = true,
  },
  [sym_children] = {
    .visible = true,
    .named = true,
  },
  [aux_sym_program_repeat1] = {
    .visible = false,
    .named = false,
  },
  [aux_sym_component_declaration_repeat1] = {
    .visible = false,
    .named = false,
  },
};

enum {
  field_body = 1,
  field_children = 2,
  field_name = 3,
};

static const char * const ts_field_names[] = {
  [0] = NULL,
  [field_body] = "body",
  [field_children] = "children",
  [field_name] = "name",
};

static const TSFieldMapSlice ts_field_map_slices[PRODUCTION_ID_COUNT] = {
  [1] = {.index = 0, .length = 1},
  [2] = {.index = 1, .length = 2},
  [3] = {.index = 3, .length = 2},
};

static const TSFieldMapEntry ts_field_map_entries[] = {
  [0] =
    {field_name, 0},
  [1] =
    {field_children, 1},
    {field_name, 0},
  [3] =
    {field_body, 3},
    {field_name, 1},
};

static const TSSymbol ts_alias_sequences[PRODUCTION_ID_COUNT][MAX_ALIAS_SEQUENCE_LENGTH] = {
  [0] = {0},
};

static const uint16_t ts_non_terminal_alias_map[] = {
  0,
};

static const TSStateId ts_primary_state_ids[STATE_COUNT] = {
  [0] = 0,
  [1] = 1,
  [2] = 2,
  [3] = 3,
  [4] = 4,
  [5] = 5,
  [6] = 6,
  [7] = 7,
  [8] = 8,
  [9] = 9,
  [10] = 10,
  [11] = 11,
  [12] = 12,
  [13] = 13,
  [14] = 14,
  [15] = 15,
};

static bool ts_lex(TSLexer *lexer, TSStateId state) {
  START_LEXER();
  eof = lexer->eof(lexer);
  switch (state) {
    case 0:
      if (eof) ADVANCE(10);
      if (lookahead == 'c') ADVANCE(5);
      if (lookahead == '\t' ||
          lookahead == '\n' ||
          lookahead == '\r' ||
          lookahead == ' ') SKIP(0)
      END_STATE();
    case 1:
      if (lookahead == 'e') ADVANCE(4);
      END_STATE();
    case 2:
      if (lookahead == 'm') ADVANCE(7);
      END_STATE();
    case 3:
      if (lookahead == 'n') ADVANCE(1);
      END_STATE();
    case 4:
      if (lookahead == 'n') ADVANCE(8);
      END_STATE();
    case 5:
      if (lookahead == 'o') ADVANCE(2);
      END_STATE();
    case 6:
      if (lookahead == 'o') ADVANCE(3);
      END_STATE();
    case 7:
      if (lookahead == 'p') ADVANCE(6);
      END_STATE();
    case 8:
      if (lookahead == 't') ADVANCE(11);
      END_STATE();
    case 9:
      if (lookahead == '\t' ||
          lookahead == '\n' ||
          lookahead == '\r' ||
          lookahead == ' ') SKIP(9)
      if (('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(12);
      END_STATE();
    case 10:
      ACCEPT_TOKEN(ts_builtin_sym_end);
      END_STATE();
    case 11:
      ACCEPT_TOKEN(anon_sym_component);
      END_STATE();
    case 12:
      ACCEPT_TOKEN(sym_identifier);
      if (('0' <= lookahead && lookahead <= '9') ||
          ('A' <= lookahead && lookahead <= 'Z') ||
          lookahead == '_' ||
          ('a' <= lookahead && lookahead <= 'z')) ADVANCE(12);
      END_STATE();
    default:
      return false;
  }
}

static const TSLexMode ts_lex_modes[STATE_COUNT] = {
  [0] = {.lex_state = 0, .external_lex_state = 1},
  [1] = {.lex_state = 0},
  [2] = {.lex_state = 0},
  [3] = {.lex_state = 0},
  [4] = {.lex_state = 9, .external_lex_state = 2},
  [5] = {.lex_state = 9, .external_lex_state = 3},
  [6] = {.lex_state = 9, .external_lex_state = 3},
  [7] = {.lex_state = 9, .external_lex_state = 3},
  [8] = {.lex_state = 9},
  [9] = {.lex_state = 9},
  [10] = {.lex_state = 9, .external_lex_state = 3},
  [11] = {.lex_state = 0},
  [12] = {.lex_state = 9, .external_lex_state = 3},
  [13] = {.lex_state = 9},
  [14] = {.lex_state = 0},
  [15] = {.lex_state = 0, .external_lex_state = 4},
};

enum {
  ts_external_token__newline = 0,
  ts_external_token__indent = 1,
  ts_external_token__dedent = 2,
};

static const TSSymbol ts_external_scanner_symbol_map[EXTERNAL_TOKEN_COUNT] = {
  [ts_external_token__newline] = sym__newline,
  [ts_external_token__indent] = sym__indent,
  [ts_external_token__dedent] = sym__dedent,
};

static const bool ts_external_scanner_states[5][EXTERNAL_TOKEN_COUNT] = {
  [1] = {
    [ts_external_token__newline] = true,
    [ts_external_token__indent] = true,
    [ts_external_token__dedent] = true,
  },
  [2] = {
    [ts_external_token__indent] = true,
    [ts_external_token__dedent] = true,
  },
  [3] = {
    [ts_external_token__dedent] = true,
  },
  [4] = {
    [ts_external_token__indent] = true,
  },
};

static const uint16_t ts_parse_table[LARGE_STATE_COUNT][SYMBOL_COUNT] = {
  [0] = {
    [ts_builtin_sym_end] = ACTIONS(1),
    [anon_sym_component] = ACTIONS(1),
    [sym__newline] = ACTIONS(1),
    [sym__indent] = ACTIONS(1),
    [sym__dedent] = ACTIONS(1),
  },
  [1] = {
    [sym_program] = STATE(14),
    [sym__fragment] = STATE(2),
    [sym_component_declaration] = STATE(2),
    [aux_sym_program_repeat1] = STATE(2),
    [ts_builtin_sym_end] = ACTIONS(3),
    [anon_sym_component] = ACTIONS(5),
  },
};

static const uint16_t ts_small_parse_table[] = {
  [0] = 3,
    ACTIONS(5), 1,
      anon_sym_component,
    ACTIONS(7), 1,
      ts_builtin_sym_end,
    STATE(3), 3,
      sym__fragment,
      sym_component_declaration,
      aux_sym_program_repeat1,
  [12] = 3,
    ACTIONS(9), 1,
      ts_builtin_sym_end,
    ACTIONS(11), 1,
      anon_sym_component,
    STATE(3), 3,
      sym__fragment,
      sym_component_declaration,
      aux_sym_program_repeat1,
  [24] = 3,
    ACTIONS(16), 1,
      sym__indent,
    STATE(10), 1,
      sym_children,
    ACTIONS(14), 2,
      sym__dedent,
      sym_identifier,
  [35] = 3,
    ACTIONS(18), 1,
      sym_identifier,
    ACTIONS(20), 1,
      sym__dedent,
    STATE(6), 2,
      sym_tag,
      aux_sym_component_declaration_repeat1,
  [46] = 3,
    ACTIONS(22), 1,
      sym_identifier,
    ACTIONS(25), 1,
      sym__dedent,
    STATE(6), 2,
      sym_tag,
      aux_sym_component_declaration_repeat1,
  [57] = 3,
    ACTIONS(18), 1,
      sym_identifier,
    ACTIONS(27), 1,
      sym__dedent,
    STATE(6), 2,
      sym_tag,
      aux_sym_component_declaration_repeat1,
  [68] = 2,
    ACTIONS(18), 1,
      sym_identifier,
    STATE(5), 2,
      sym_tag,
      aux_sym_component_declaration_repeat1,
  [76] = 2,
    ACTIONS(18), 1,
      sym_identifier,
    STATE(7), 2,
      sym_tag,
      aux_sym_component_declaration_repeat1,
  [84] = 1,
    ACTIONS(29), 2,
      sym__dedent,
      sym_identifier,
  [89] = 1,
    ACTIONS(31), 2,
      ts_builtin_sym_end,
      anon_sym_component,
  [94] = 1,
    ACTIONS(33), 2,
      sym__dedent,
      sym_identifier,
  [99] = 1,
    ACTIONS(35), 1,
      sym_identifier,
  [103] = 1,
    ACTIONS(37), 1,
      ts_builtin_sym_end,
  [107] = 1,
    ACTIONS(39), 1,
      sym__indent,
};

static const uint32_t ts_small_parse_table_map[] = {
  [SMALL_STATE(2)] = 0,
  [SMALL_STATE(3)] = 12,
  [SMALL_STATE(4)] = 24,
  [SMALL_STATE(5)] = 35,
  [SMALL_STATE(6)] = 46,
  [SMALL_STATE(7)] = 57,
  [SMALL_STATE(8)] = 68,
  [SMALL_STATE(9)] = 76,
  [SMALL_STATE(10)] = 84,
  [SMALL_STATE(11)] = 89,
  [SMALL_STATE(12)] = 94,
  [SMALL_STATE(13)] = 99,
  [SMALL_STATE(14)] = 103,
  [SMALL_STATE(15)] = 107,
};

static const TSParseActionEntry ts_parse_actions[] = {
  [0] = {.entry = {.count = 0, .reusable = false}},
  [1] = {.entry = {.count = 1, .reusable = false}}, RECOVER(),
  [3] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_program, 0),
  [5] = {.entry = {.count = 1, .reusable = true}}, SHIFT(13),
  [7] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_program, 1),
  [9] = {.entry = {.count = 1, .reusable = true}}, REDUCE(aux_sym_program_repeat1, 2),
  [11] = {.entry = {.count = 2, .reusable = true}}, REDUCE(aux_sym_program_repeat1, 2), SHIFT_REPEAT(13),
  [14] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_tag, 1, .production_id = 1),
  [16] = {.entry = {.count = 1, .reusable = true}}, SHIFT(9),
  [18] = {.entry = {.count = 1, .reusable = true}}, SHIFT(4),
  [20] = {.entry = {.count = 1, .reusable = true}}, SHIFT(11),
  [22] = {.entry = {.count = 2, .reusable = true}}, REDUCE(aux_sym_component_declaration_repeat1, 2), SHIFT_REPEAT(4),
  [25] = {.entry = {.count = 1, .reusable = true}}, REDUCE(aux_sym_component_declaration_repeat1, 2),
  [27] = {.entry = {.count = 1, .reusable = true}}, SHIFT(12),
  [29] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_tag, 2, .production_id = 2),
  [31] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_component_declaration, 5, .production_id = 3),
  [33] = {.entry = {.count = 1, .reusable = true}}, REDUCE(sym_children, 3),
  [35] = {.entry = {.count = 1, .reusable = true}}, SHIFT(15),
  [37] = {.entry = {.count = 1, .reusable = true}},  ACCEPT_INPUT(),
  [39] = {.entry = {.count = 1, .reusable = true}}, SHIFT(8),
};

#ifdef __cplusplus
extern "C" {
#endif
void *tree_sitter_quanty_external_scanner_create(void);
void tree_sitter_quanty_external_scanner_destroy(void *);
bool tree_sitter_quanty_external_scanner_scan(void *, TSLexer *, const bool *);
unsigned tree_sitter_quanty_external_scanner_serialize(void *, char *);
void tree_sitter_quanty_external_scanner_deserialize(void *, const char *, unsigned);

#ifdef _WIN32
#define extern __declspec(dllexport)
#endif

extern const TSLanguage *tree_sitter_quanty(void) {
  static const TSLanguage language = {
    .version = LANGUAGE_VERSION,
    .symbol_count = SYMBOL_COUNT,
    .alias_count = ALIAS_COUNT,
    .token_count = TOKEN_COUNT,
    .external_token_count = EXTERNAL_TOKEN_COUNT,
    .state_count = STATE_COUNT,
    .large_state_count = LARGE_STATE_COUNT,
    .production_id_count = PRODUCTION_ID_COUNT,
    .field_count = FIELD_COUNT,
    .max_alias_sequence_length = MAX_ALIAS_SEQUENCE_LENGTH,
    .parse_table = &ts_parse_table[0][0],
    .small_parse_table = ts_small_parse_table,
    .small_parse_table_map = ts_small_parse_table_map,
    .parse_actions = ts_parse_actions,
    .symbol_names = ts_symbol_names,
    .field_names = ts_field_names,
    .field_map_slices = ts_field_map_slices,
    .field_map_entries = ts_field_map_entries,
    .symbol_metadata = ts_symbol_metadata,
    .public_symbol_map = ts_symbol_map,
    .alias_map = ts_non_terminal_alias_map,
    .alias_sequences = &ts_alias_sequences[0][0],
    .lex_modes = ts_lex_modes,
    .lex_fn = ts_lex,
    .external_scanner = {
      &ts_external_scanner_states[0][0],
      ts_external_scanner_symbol_map,
      tree_sitter_quanty_external_scanner_create,
      tree_sitter_quanty_external_scanner_destroy,
      tree_sitter_quanty_external_scanner_scan,
      tree_sitter_quanty_external_scanner_serialize,
      tree_sitter_quanty_external_scanner_deserialize,
    },
    .primary_state_ids = ts_primary_state_ids,
  };
  return &language;
}
#ifdef __cplusplus
}
#endif
