export default class WebViewService {
  private static readonly bindings: NodeJS.Global = global;

  static async add(a: number, b: number): Promise<number> {
    try {
      const result = await WebViewService.bindings.add(a, b);
      return result;
    } catch (error) {
      throw error;
    }
  }

  static async printObject(obj: Object): Promise<void> {
    try {
      await WebViewService.bindings.printObject(obj);
    } catch (error) {
      throw error;
    }
  }
}
